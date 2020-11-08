package controller

import (
	"log"
	"net/http"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/services"
	"github.com/gin-gonic/gin"
)

func DeleteBookController(ctx *gin.Context) {
	log.Println("#DeleteBookController start;")

	var request DeleteBookComp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Requested with invalid parameters;")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Invalid Parameters": err.Error()})
		return
	}

	err = DeleteBookService(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	} else {
		ctx.JSONP(http.StatusOK, gin.H{
			"message":  "ok",
			"response": "delete succeded",
		})
	}

	log.Println("#DeleteBookController end;")
}
