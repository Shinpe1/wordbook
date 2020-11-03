package controller

import (
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳を新規登録します */
func RegisterBookController(ctx *gin.Context) {
	// パラメータの取得
	var request NewBook
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameters": err.Error()})
	}

	// 新しくレコードをbooksテーブルに挿入します
	RegisterNewBookService(&request)

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": "insert succeeded",
	})
}
