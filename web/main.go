package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ryo246912/playground-go-web/query"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func setup() *bun.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbDatabase = os.Getenv("DB_NAME")
		dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
			dbPassword, dbDatabase)
	)

	sqldb, err := sql.Open("mysql", dbConn)
	if err != nil {
		panic("error")
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}

func main() {
	if len(os.Args) >= 1 {
		db := setup()
		switch os.Args[1] {
		case "1":
			query.Query1(db)
		case "2":
			fmt.Println("2を受け取りました")
			// ここに2のときの処理を書く
		default:
			fmt.Println("不明な引数:", os.Args[1])
		}
		return
	}

	fmt.Println("引数がありません")
}
