package repositories

import (
	"database/sql"

	"github.com/ryo246912/path-to-intermediate-go-developer/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
insert into articles (title, contents, username, nice, created_at) values
(?, ?, ?, 0, now());
`
	// (問 1) 構造体 `models.Article`を受け取って、それをデータベースに挿入する処理
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}
	newArticle := models.Article{
		ID:       int(id),
		Title:    article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
		NiceNum:  0,
	}

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
select article_id, title, contents, username, nice
from articles
limit ? offset ?;
`
	// (問 2) 指定された記事データをデータベースから取得して、
	// それを models.Article 構造体のスライス []models.Article に詰めて返す処理
	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articleArray []models.Article
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
select *
from articles
where article_id = ?;
`
	// (問 3) 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の形で返す処理↪
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
select nice
from articles
where article_id = ?;
`
	// (問 4) 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err = row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var nice int
	err = row.Scan(&nice)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = db.Exec(sqlUpdateNice, nice+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
