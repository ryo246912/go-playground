package services

import (
	"github.com/ryo246912/path-to-intermediate-go-developer/models"
	"github.com/ryo246912/path-to-intermediate-go-developer/repositories"
)

// PostCommentHandler で使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	return repositories.InsertComment(s.db, comment)
}
