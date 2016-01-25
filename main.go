package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	logger      *log.Logger
	goPath      = os.Getenv("GOPATH")
	gocodePath  = goPath + "/bin/gocode"
	projectPath = goPath + "/src/github.com/gocode"
)

func init() {
	logger = log.New(os.Stdout, "epubviewer ", log.Lshortfile)
}

func staticFilesHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, r.URL.Path[1:])
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "static/editor.html")
}

func dirHandler(rw http.ResponseWriter, r *http.Request) {
	dir := ReadDir(projectPath)

	enc := json.NewEncoder(rw)
	if err := enc.Encode(dir); err != nil {
		logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func goFileHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		buf, err := ioutil.ReadFile(goPath + req.URL.Path)
		if err != nil {
			logger.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Write(buf)
	} else {
		homeHandler(rw, req)
	}
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	content := req.FormValue("content")

	var formatedContent []byte
	if strings.HasSuffix(strings.ToLower(name), ".go") {
		var err error
		formatedContent, err = formatSource([]byte(content))
		if err != nil {
			logger.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		formatedContent = []byte(content)
	}

	err := ioutil.WriteFile(goPath+name, formatedContent, os.ModePerm)
	if err != nil {
		logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.Write(formatedContent)
}

func formatHandler(rw http.ResponseWriter, req *http.Request) {
	content := req.FormValue("content")

	result, err := formatSource([]byte(content))
	if err != nil {
		logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(result)
}

func autocompleteHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	content := req.FormValue("content")
	offset := req.FormValue("offset")

	err := ioutil.WriteFile(goPath+name, []byte(content), os.ModePerm)
	if err != nil {
		logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := autoComplete(name, []byte(content), offset)

	buf, _ := json.Marshal(result)
	rw.Write(buf)
}

func initHandler(rw http.ResponseWriter, r *http.Request) {
	type response struct {
		GoPath      string `json:"gopath"`
		GocodePath  string `json:"gocodePath"`
		ProjectPath string `json:"projectPath"`
	}

	resp := &response{GoPath: goPath, GocodePath: gocodePath, ProjectPath: projectPath}

	enc := json.NewEncoder(rw)
	if err := enc.Encode(resp); err != nil {
		logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func saveSettingsHandler(rw http.ResponseWriter, req *http.Request) {
	goPath = req.FormValue("gopath")
	gocodePath = req.FormValue("gocodePath")
	projectPath = req.FormValue("projectPath")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dir", dirHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/src/", goFileHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/format", formatHandler)
	http.HandleFunc("/saveSettings", saveSettingsHandler)
	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/autocomplete", autocompleteHandler)
	http.ListenAndServe(":9090", nil)
}
