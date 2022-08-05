package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func generateData() map[string]map[string]string {
	vars := map[string]map[string]string{
		"envs": make(map[string]string),
	}

	var envDirs string
	envDirs = os.Getenv("SECRETS_DIRS")
	if envDirs != "" {
		var dirs []string
		dirs = strings.Split(envDirs, ",")
		fmt.Println(dirs)

		var files []string

		for idx := range dirs {
			err := filepath.Walk(dirs[idx], func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return nil
			})
			if err != nil {
				fmt.Printf("An error occured: %s\n", err)
			}
		}

		for idx := range files {
			pair, err := readVarFromFile(files[idx])
			if err != nil {
				fmt.Printf("An error occured: %s\n", err)
			}

			if pair != nil {
				_, ok := vars["file"]
				if ok != true {
					vars["file"] = map[string]string{
						pair["key"]: pair["val"],
					}
				} else {
					vars["file"][pair["key"]] = pair["val"]
				}
			}
		}
	}

	for _, envs := range os.Environ() {
		pair := strings.SplitN(envs, "=", 2)
		vars["envs"][pair[0]] = pair[1]
	}

	return vars
}

func readVarFromFile(filename string) (map[string]string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return nil, errors.New("something went wrong with access to file")
	}
	if info.IsDir() {
		return nil, errors.New("it should be file, not a directory")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("can't read from file")
	}

	pair := make(map[string]string)
	pair["key"] = strings.ToUpper(filepath.Base(filename))
	pair["val"] = string(data)
	return pair, nil
}

func renderWeb(w http.ResponseWriter, r *http.Request) {
	// TODO: should return json
	vars := generateData()

	renderedPage, _ := template.ParseFiles("envs.gohtml")
	err := renderedPage.Execute(w, vars)
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	}
	// TODO: use a normal logger
	logMessage := map[string]string{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}
	logMsgJSON, marshallErr := json.Marshal(logMessage)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	fmt.Println(string(logMsgJSON))
}

func ping(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	jsonData, marshallErr := json.Marshal(data)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	// TODO: use a normal logger
	logMessage := map[string]string{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}
	logMsgJSON, marshallErr := json.Marshal(logMessage)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	fmt.Println(string(logMsgJSON))
}

func jsonEnvs(w http.ResponseWriter, r *http.Request) {
	data := generateData()

	jsonData, marshallErr := json.Marshal(data)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	// TODO: use a normal logger
	logMessage := map[string]string{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}
	logMsgJSON, marshallErr := json.Marshal(logMessage)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	fmt.Println(string(logMsgJSON))
}

func checkServices(w http.ResponseWriter, r *http.Request) {
	type host struct {
		Addr string
		Code string
	}

	type data struct {
		Hosts []host
		Error string
	}

	var response data

	hostsList := os.Getenv("HOSTS")
	if hostsList == "" {
		response.Error = "Variable HOSTS is empty, should be list"
	} else {
		hosts := strings.Split(hostsList, ";")

		for _, hostAddr := range hosts {
			code := "0"
			resp, respErr := http.Get(hostAddr)
			if respErr != nil {
				fmt.Printf("An error occured: %s\n", respErr)
			} else {
				code = strconv.Itoa(resp.StatusCode)
			}

			response.Hosts = append(response.Hosts, host{hostAddr, code})

			// TODO: use a normal logger
			logMessage := map[string]string{
				"Host": hostAddr,
				"Code": code,
			}
			logMsgJSON, marshallErr := json.Marshal(logMessage)
			if marshallErr != nil {
				fmt.Printf("An error occured: %s\n", marshallErr)
			}
			fmt.Println(string(logMsgJSON))
		}
	}

	jsonData, dataMarshallErr := json.Marshal(response)
	if dataMarshallErr != nil {
		fmt.Printf("An error occured: %s\n", dataMarshallErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	// TODO: use a normal logger
	logMessage := map[string]string{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}
	logMsgJSON, marshallErr := json.Marshal(logMessage)
	if marshallErr != nil {
		fmt.Printf("An error occured: %s\n", marshallErr)
	}
	fmt.Println(string(logMsgJSON))
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8080"
	}

	fmt.Printf("Server is starting on: %v\n", listenAddr)
	http.HandleFunc("/", renderWeb)
	http.HandleFunc("/json", jsonEnvs)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/net-check", checkServices)
	_ = http.ListenAndServe(listenAddr, nil)
}
