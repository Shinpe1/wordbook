package endpoint

import (
	"log"
	"net/http"
	"strconv"

	services "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳一覧取得 */
func GetBookListController(ctx *gin.Context) {
	log.Println("#GetBookListController start;")

	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter"})
		return
	}

	books, err := services.GetListService(userId)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{"message": "request success but no content returned"})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": books,
	})

	log.Println("#GetBookListController end;")
}

/** 個別単語帳取得 */
func GetIndividualBookController(ctx *gin.Context) {
	log.Println("#GetIndividualBookController start;")

	// userIdをint型にパース
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter"})
		return
	}
	// bookIdを基数10の64bit型に変換
	bookId, err := strconv.ParseInt(ctx.Query("bookId"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter"})
		return
	}

	contents, err := services.GetIndividualBookService(userId, bookId)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{"message": "request success but no content returned"})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": contents,
	})

	log.Println("#GetIndividualBookController end;")

}
