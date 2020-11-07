package controller

import (
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

func InsertBookController(ctx *gin.Context) {
	var request InsertBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// 400エラー
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameters": err.Error()})
	}

	err = InsertBookService(&request)
	if err != nil {
		// 503エラー（DB更新エラーでこれを返していいのか？）
		ctx.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"Couldn't save records into Database;": err.Error()})
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": "insert succeeded",
	})

}
