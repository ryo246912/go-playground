package main

import (
	"log"
	"net/http"

	"github.com/ryo246912/path-to-intermediate-go-developer/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.Hellohandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.GetArticleListHandler)
	http.HandleFunc("/article/1", handlers.GetArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceArticleHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
