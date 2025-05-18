package service

import (
	"github.com/ryo246912/path-to-intermediate-go-developer/models"
	"github.com/ryo246912/path-to-intermediate-go-developer/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)
	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	return repositories.InsertArticle(db, article)
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	for i, article := range articles {
		var comments []models.Comment
		comments, _ = repositories.SelectCommentList(db, article.ID)
		articles[i].CommentList = comments
	}

	return articles, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	repositories.UpdateNiceNum(db, article.ID)

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
