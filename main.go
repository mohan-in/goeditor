package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	logger *log.Logger
	GoPath = "C:/wrk"
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
	dir := ReadDir(r.FormValue("dir"))

	enc := json.NewEncoder(rw)
	if err := enc.Encode(dir); err != nil {
		fmt.Println(err)
	}
}

func goFileHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		buf, err := ioutil.ReadFile(GoPath + req.URL.Path)
		if err != nil {
			logger.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}

		rw.Write(buf)
	} else {
		homeHandler(rw, req)
	}
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	content := req.FormValue("content")

	err := ioutil.WriteFile(GoPath+name, []byte(content), os.ModePerm)
	if err != nil {
		logger.Println(err)
	}
}

func autocompleteHandler(rw http.ResponseWriter, req *http.Request) {
	content := req.FormValue("content")
	offset := req.FormValue("offset")

	result := autoComplete([]byte(content), offset)

	for _, s := range result {
		logger.Println(s)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dir", dirHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/src/", goFileHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/autocomplete", autocompleteHandler)
	http.ListenAndServe(":9090", nil)
}
