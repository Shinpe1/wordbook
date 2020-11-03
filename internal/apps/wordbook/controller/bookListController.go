package controller

import (
	"net/http"
	"strconv"

	services "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳一覧取得 */
func GetBookListController(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameter": err.Error()})
	}

	books := services.GetListService(userId)

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": books,
	})
}

/** 個別単語帳取得 */
func GetIndividualBookController(ctx *gin.Context) {

	// userIdをint型にパース
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameter": err.Error()})
	}
	// bookIdを基数10の64bit型に変換
	bookId, err := strconv.ParseInt(ctx.Query("bookId"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameter": err.Error()})
	}

	contents := services.GetIndividualBookService(userId, bookId)

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": contents,
	})

}
