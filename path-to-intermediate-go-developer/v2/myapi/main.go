package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryo246912/path-to-intermediate-go-developer/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.Hellohandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// http.ListenAndServe 関数の第二引数というのは、実は「サーバーの中で使うルータを指定する」部分
	// ここにルータが渡されず nil だった場合には、Go の HTTP サーバーがデフォルトで持っているルータが自動的に採用されます
	log.Fatal(http.ListenAndServe(":8080", r))
}
