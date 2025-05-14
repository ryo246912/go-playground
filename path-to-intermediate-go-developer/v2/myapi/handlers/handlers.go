package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ryo246912/path-to-intermediate-go-developer/models"
)

func Hellohandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	// ストリームから直接リクエストデータを取るようにしたことで、
	// デコード前の「Content-Lengthヘッダフィールドの値からバイトスライスを作り、そこにリクエストボディの中身を書き込む」という操作がいらないなっています。
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
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

	articles := []models.Article{
		models.Article1,
		models.Article2,
	}
	json.NewEncoder(w).Encode(articles)
	io.WriteString(w, fmt.Sprintf("Article List (page %d)\n", page))
}
func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invaild query parameter", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(models.Article1)
	io.WriteString(w, fmt.Sprintf("Article No%d\n", articleID))
}
func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqContent any
	if err := json.NewDecoder(req.Body).Decode(&reqContent); err != nil {
		http.Error(w, "Invaild json decode", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(reqContent)
	io.WriteString(w, "Posting Nice...\n")
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqContent any
	if err := json.NewDecoder(req.Body).Decode(&reqContent); err != nil {
		http.Error(w, "Invaild json decode", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(reqContent)
	io.WriteString(w, "Posting Comment...\n")
}
