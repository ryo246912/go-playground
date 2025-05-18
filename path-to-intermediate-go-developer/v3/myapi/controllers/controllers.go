package controllers

import (
	"github.com/ryo246912/path-to-intermediate-go-developer/controllers/services"
)

type MyAppController struct {
	// MyAppService 構造体だった場所を MyAppServicer インターフェースに置き換え
	// これまでservice packageの内容に依存していた(services.MyAppServiceが必ず必要)が、
	// 必要だったのはMyAppService 構造体が持つメソッドを使いたかったからだけであるため、インターフェイスによる抽象化
	// service *services.MyAppService
	service services.MyAppServicer
}

func NewMyAppController(s services.MyAppServicer) *MyAppController {
	return &MyAppController{service: s}
}
