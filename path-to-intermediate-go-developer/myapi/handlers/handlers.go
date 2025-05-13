package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Hellohandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}
func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	io.WriteString(w, fmt.Sprintf("Article No.%d\n", articleID))
}
func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
