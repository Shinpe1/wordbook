package controller

import (
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳の更新を行います */
func UpdateBookController(ctx *gin.Context) {
	// リクエストボディの取得
	var request UpdateBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameters": err.Error()})
	}

	// 指定された単語帳の内容を更新します
	UpdateBookService(&request)

}
