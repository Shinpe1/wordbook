package controller

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

func InsertBookController(ctx *gin.Context) {
	log.Println("#InsertBookController start;")

	var request InsertBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// 400エラー
		log.Println("Requested with invalid parameters")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameters": err.Error()})
		return
	}

	err = InsertBookService(&request)
	if err != nil {
		log.Println(err.Error())
		// 503エラー（DB更新エラーでこれを返していいのか？）
		ctx.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"Couldn't save records into Database;": err.Error()})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": "insert succeeded",
	})

	log.Println("#InsertBookController end;")
}
