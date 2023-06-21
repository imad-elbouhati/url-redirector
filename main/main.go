package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"urlredirect"
)

const (
	YAML_PATH_FILE = "./yaml-data.yaml"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlredirect.MapHandler(pathsToUrls, mux)

	
	yaml := readYamlFile(YAML_PATH_FILE)

	yamlHandler, err := urlredirect.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readYamlFile(path string) string {
	content,err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Print(err)
	}

	return string(content)
}
