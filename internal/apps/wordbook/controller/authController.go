package controller

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/response"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

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

func LoginUserController(ctx *gin.Context) {
	log.Println("#LoginUserController start;")

	var request LoginUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter"})
		return
	}

	var response LoginUserResponse
	response, err = LoginUserService(request)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message":  "ok",
		"response": response,
	})

	log.Println("#LoginUserController end;")
}

func RefreshTokenController(ctx *gin.Context) {

}
