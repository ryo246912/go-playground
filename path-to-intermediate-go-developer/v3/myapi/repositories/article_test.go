package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ryo246912/path-to-intermediate-go-developer/models"
	"github.com/ryo246912/path-to-intermediate-go-developer/repositories"
	testdata "github.com/ryo246912/path-to-intermediate-go-developer/repositories/test_data"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title,
					test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents,
					test.expected.Contents)
			}
		},
		)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}
	// expectedArticleNum := 3
	expectTitle := "insertTest"
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	// if newArticle.Title != expectedArticleNum {
	// 	t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum,
	// 		newArticle.ID)
	// }
	if newArticle.Title != expectTitle {
		t.Errorf("new article title is expected %s but got %s\n", expectTitle, "insertTest")
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from articles
		where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})

}

func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	setup() // 前処理
	m.Run()
	teardown() // 後処理
}
