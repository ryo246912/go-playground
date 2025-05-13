package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Hellohandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invaild query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	io.WriteString(w, fmt.Sprintf("Article List (page %d)\n", page))
}
func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invaild query parameter", http.StatusBadRequest)
	}
	io.WriteString(w, fmt.Sprintf("Article No.%d\n", articleID))
}
func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
