package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ReadVarFromFile(filename string) (map[string]string, error) {
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

func ShowVars(w http.ResponseWriter, r *http.Request) {
	vars := map[string]map[string]string{
		"envs": make(map[string]string),
	}

  var env_dirs string
  env_dirs = os.Getenv("SECRETS_DIRS")
	var dirs []string
	dirs = strings.Split(env_dirs, ",")
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

	for _, envs := range os.Environ() {
		pair := strings.SplitN(envs, "=", 2)
		vars["envs"][pair[0]] = pair[1]
	}

	for idx := range files {
		pair, err := ReadVarFromFile(files[idx])
		if err != nil {
			fmt.Printf("An error occured: %s\n", err)
		}

		if pair != nil {
			_, ok := vars["file"]
			if ok != true {
				vars["file"] = map[string]string {
					pair["key"]: pair["val"],
				}
			} else {
				vars["file"][pair["key"]] = pair["val"]
			}
		}
	}

	renderedPage, _ := template.ParseFiles("envs.gohtml")
	err := renderedPage.Execute(w, vars)
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	}
	fmt.Printf("{\"Method\": \"%s\", \"url\": \"%s\", \"Remote-Addr\":\"%s\"}\n", r.Method, r.URL, r.RemoteAddr)
}

func main() {
	fmt.Println("Server is starting on: 0.0.0.0:8080")
	http.HandleFunc("/", ShowVars)
	_ = http.ListenAndServe(":8080", nil)
}
