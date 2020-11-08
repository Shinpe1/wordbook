package controller

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳の更新を行います */
func UpdateBookController(ctx *gin.Context) {
	log.Println("#UpdateBookController start;")
	// リクエストボディの取得
	var request UpdateBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	// 指定された単語帳の内容を更新します
	err = UpdateBookService(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Couldn't update records"})
		return
	}

	log.Println("#UpdateBookController end;")
}
