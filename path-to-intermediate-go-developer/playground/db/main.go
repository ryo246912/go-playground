package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	models "github.com/ryo246912/path-to-intermediate-go-developer/model"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 1
	// rows, err := db.Query(`select * from articles where article_id = ?;`, articleID)
	row := db.QueryRow(`select * from articles where article_id = ?;`, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}
	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)

	// 新規作成処理
	// newArticle := models.Article{
	// 	Title:    "insert test",
	// 	Contents: "Can I insert data correctly?",
	// 	UserName: "saki",
	// }
	// result, err := db.Exec(`insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());`, newArticle.Title, newArticle.Contents, newArticle.UserName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	// 更新処理
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	niceRow := tx.QueryRow(`select nice from articles where article_id = ?;`, articleID)
	if err := niceRow.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var nicenum int
	err = niceRow.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	fmt.Println(nicenum)
	_, err = tx.Exec(`update articles set nice = ? where article_id = ?`, nicenum+1, articleID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	tx.Commit()

}
