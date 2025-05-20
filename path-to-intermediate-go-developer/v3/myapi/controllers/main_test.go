package controllers_test

import (
	"testing"

	"github.com/ryo246912/path-to-intermediate-go-developer/controllers"
	"github.com/ryo246912/path-to-intermediate-go-developer/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()

	aCon = controllers.NewArticleController(ser)
	m.Run()
}
