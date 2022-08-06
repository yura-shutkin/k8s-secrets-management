package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
		log.WithFields(log.Fields{
			"dirs": dirs,
		}).Info()

		var files []string

		for idx := range dirs {
			err := filepath.Walk(dirs[idx], func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return nil
			})
			if err != nil {
				log.WithFields(log.Fields{
					"Error": err,
				}).Error()
			}
		}

		for idx := range files {
			pair := readVarFromFile(files[idx])

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

func readVarFromFile(filename string) map[string]string {
	info, err := os.Stat(filename)
	if err != nil {
		log.Errorf("Something went wrong with access to file %v", filename)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Errorf("Can't read from file %v", filename)
		return nil
	}

	pair := make(map[string]string)
	pair["key"] = strings.ToUpper(filepath.Base(filename))
	pair["val"] = string(data)
	return pair
}

func renderHtml(w http.ResponseWriter, r *http.Request) {
	vars := generateData()

	renderedPage, _ := template.ParseFiles("envs.gohtml")
	err := renderedPage.Execute(w, vars)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error()
	}

	log.WithFields(log.Fields{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}).Info()
}

func ping(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	jsonData, marshallErr := json.Marshal(data)
	if marshallErr != nil {
		log.WithFields(log.Fields{
			"Error": marshallErr,
		}).Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	log.WithFields(log.Fields{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}).Info()
}

func jsonEnvs(w http.ResponseWriter, r *http.Request) {
	data := generateData()

	jsonData, marshallErr := json.Marshal(data)
	if marshallErr != nil {
		log.WithFields(log.Fields{
			"Error": marshallErr,
		}).Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	log.WithFields(log.Fields{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}).Info()
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
				log.WithFields(log.Fields{
					"Error": respErr,
				}).Error()
			} else {
				code = strconv.Itoa(resp.StatusCode)
			}

			response.Hosts = append(response.Hosts, host{hostAddr, code})

			log.WithFields(log.Fields{
				"Host": hostAddr,
				"Code": code,
			}).Info()
		}
	}

	jsonData, dataMarshallErr := json.Marshal(response)
	if dataMarshallErr != nil {
		log.WithFields(log.Fields{
			"Error": dataMarshallErr,
		}).Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	log.WithFields(log.Fields{
		"Method":      r.Method,
		"URL":         r.URL.String(),
		"Remote-Addr": r.RemoteAddr,
	}).Info()
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "localhost:8080"
	}

	log.Infof("Server is starting on: %v", listenAddr)
	http.HandleFunc("/", renderHtml)
	http.HandleFunc("/json", jsonEnvs)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/net-check", checkServices)
	_ = http.ListenAndServe(listenAddr, nil)
}
