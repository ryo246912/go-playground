package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"create_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {

	article := Article{
		ID:       1,
		Title:    "タイトル",
		Contents: "記事の中身",
		UserName: "鈴木太郎",
		NiceNum:  1,
		CommentList: []Comment{
			{
				CommentID: 1,
				ArticleID: 1,
				Message:   "テスト1",
				CreatedAt: time.Now(),
			},
		},
		CreatedAt: time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
