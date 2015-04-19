package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dir", dirHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.ListenAndServe(":8080", nil)
}

func dirHandler(rw http.ResponseWriter, r *http.Request) {

	dir := ReadDir(r.FormValue("dir"))

	enc := json.NewEncoder(rw)
	if err := enc.Encode(dir); err != nil {
		fmt.Println(err)
	}
}

func staticFilesHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, r.URL.Path[1:])
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "static/editor.html")
}