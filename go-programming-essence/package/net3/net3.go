package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	mimemap := map[string]string{
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".ppt":  "application/vnd.ms-powerpoint",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		".pdf":  "application/pdf",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if typ, ok := mimemap[path.Ext(r.URL.Path)]; ok {
			w.Header().Set("Content-Type", typ)
		}
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
