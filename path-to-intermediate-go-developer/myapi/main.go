package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	hellohandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}
	getArticleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}
	getArticleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		io.WriteString(w, fmt.Sprintf("Article No.%d\n", articleID))
	}
	postNiceArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}
	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", getArticleListHandler)
	http.HandleFunc("/article/1", getArticleDetailHandler)
	http.HandleFunc("/article/nice", postNiceArticleHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
