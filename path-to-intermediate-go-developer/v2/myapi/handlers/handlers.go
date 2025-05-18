package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ryo246912/path-to-intermediate-go-developer/models"
	"github.com/ryo246912/path-to-intermediate-go-developer/service"
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

	article, err := service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
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

	articles, err := service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invaild query parameter", http.StatusBadRequest)
	}

	article, err := service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Invaild json decode", http.StatusBadRequest)
		return
	}

	article, err := service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(article)
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Invaild json decode", http.StatusBadRequest)
		return
	}

	comment, err := service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
