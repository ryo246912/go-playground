package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ryo246912/path-to-intermediate-go-developer/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbDatabase = os.Getenv("DB_NAME")
		dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
			dbPassword, dbDatabase)
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		panic("error")
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	// http.ListenAndServe 関数の第二引数というのは、実は「サーバーの中で使うルータを指定する」部分
	// ここにルータが渡されず nil だった場合には、Go の HTTP サーバーがデフォルトで持っているルータが自動的に採用されます
	log.Fatal(http.ListenAndServe(":8080", r))
}
