package wordbook

import (
	"log"
	"net/http"
	"os"

	"github.com/Shinpe1/wordbook_web/internal/apps/wordbook/controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	. "github.com/Shinpe1/wordbook_web/auth"
)

func Run() {
	router := gin.Default()

	// コンテクストパス
	const CONTEXT_PATH = "/api/v1.0"
	ORIGIN_PATH := router.Group(CONTEXT_PATH)
	{
		// 取得API
		listUrl := ORIGIN_PATH.Group("/list", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳一覧取得
			listUrl.GET("/all", controller.GetBookListController)
			// 個別単語帳取得
			listUrl.GET("/individual", controller.GetIndividualBookController)
		}
		// 登録API
		registerUrl := ORIGIN_PATH.Group("/register", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳新規登録
			registerUrl.POST("", controller.RegisterBookController)
		}
		// 更新API
		updateUrl := ORIGIN_PATH.Group("/update", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳更新
			updateUrl.POST("", controller.UpdateBookController)
		}
		// 追加API
		insertUrl := ORIGIN_PATH.Group("/insert", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳追加
			insertUrl.POST("", controller.InsertBookController)
		}
		// 削除API
		deleteUrl := ORIGIN_PATH.Group("/delete", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳削除
			deleteUrl.POST("", controller.DeleteBookController)
		}
		//認証API
		authUrl := ORIGIN_PATH.Group("/auth")
		{
			// ユーザー登録
			authUrl.POST("/register", controller.RegisterUserController)
			// ログイン
			// authUrl.POST("/login", controller.LoginUserController)
			authUrl.POST("/login", AuthMiddleware.LoginHandler)
			// トークンリフレッシュ
			authUrl.POST("/refresh", AuthMiddleware.RefreshHandler)
		}
	}

	router.NoRoute(AuthMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Println("NoRoute Claims: {}", claims)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"code":    "PAGE NOT FOUND",
			"message": "page not found",
		})
	})

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err.Error())
	}
}
