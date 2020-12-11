package endpoint

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

func RegisterTempUserController(ctx *gin.Context) {
	log.Println("#RegisterTempUserController start;")

	var request TempUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	err = RegisterTempUserService(request)
	if err != nil {
		log.Fatal(err.Error())
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "Service not working tempolarly"})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": nil,
	})

	log.Println("#RegisterTempUserController end;")
}

func RegisterUserController(ctx *gin.Context) {
	log.Println("#RegisterUserController start;")

	var request RegisterUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter"})
		return
	}

	err = RegisterUserService(request)
	if err != nil {
		log.Fatal(err.Error())
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "Failed to register"})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": nil,
	})

	log.Println("#RegisterUserController end;")
}

func ValidateTempUserController(ctx *gin.Context) {
	log.Println("#ValidateTempUserController start;")

	token := ctx.Query("token")

	res, err := ValidateTempTokenService(token)
	if err != nil {
		ctx.HTML(http.StatusServiceUnavailable, "done.tmpl", gin.H{
			"title":    "登録失敗",
			"message":  "エラーが発生しました。仮登録をやり直してください",
			"response": nil,
		})
		return
	}

	if res {
		// 有効なトークンだった場合
		ctx.HTML(http.StatusOK, "done.tmpl", gin.H{
			"title":    "登録完了",
			"message":  "登録完了。アプリに戻って登録を続けてください",
			"response": nil,
		})
	} else {
		// 無効なトークンだった場合
		ctx.HTML(http.StatusBadRequest, "done.tmpl", gin.H{
			"title":    "登録失敗",
			"message":  "登録に失敗しました。有効期限を過ぎている可能性があります。もう一度仮登録を行ってください",
			"response": nil,
		})
		// ctx.Redirect(http.StatusSeeOther, "wordbook-anywhere.app://")
	}

	log.Println("#ValidateTempUserController end;")
}
