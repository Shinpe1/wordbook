package endpoint

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

/** 単語帳を新規登録します */
func RegisterBookController(ctx *gin.Context) {
	// リクエストボディの取得
	// JSONの形にキャスト
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

func DeleteBookController(ctx *gin.Context) {
	log.Println("#DeleteBookController start;")

	var request DeleteBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	err = DeleteBookService(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSONP(http.StatusOK, gin.H{
			"message":  "ok",
			"response": "delete succeded",
		})
	}

	log.Println("#DeleteBookController end;")
}
