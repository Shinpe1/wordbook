package wordbook

import (
	"log"
	"net/http"
	"os"

	"github.com/Shinpe1/wordbook_web/internal/apps/wordbook/endpoint"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	. "github.com/Shinpe1/wordbook_web/auth"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/tmpl/*")

	// コンテクストパス
	const CONTEXT_PATH = "/api/v1.0"
	ORIGIN_PATH := router.Group(CONTEXT_PATH)
	{
		// 単語帳API
		bookAPI := ORIGIN_PATH.Group("/book", AuthMiddleware.MiddlewareFunc())
		{
			getBookAPI := bookAPI.Group("/list")
			{
				// 単語帳一覧取得
				getBookAPI.GET("/all", endpoint.GetBookListController)
				// 個別単語帳取得
				getBookAPI.GET("/detail", endpoint.GetIndividualBookController)
			}

			// 単語帳新規登録
			bookAPI.POST("/register", endpoint.RegisterBookController)
			// 単語帳追加
			bookAPI.POST("/insert", endpoint.InsertBookController)
			// 単語帳更新
			bookAPI.POST("/update", endpoint.UpdateBookController)
			// 単語帳削除
			bookAPI.POST("/delete", endpoint.DeleteBookController)
		}
		// 取得API
		listUrl := ORIGIN_PATH.Group("/list", AuthMiddleware.MiddlewareFunc())
		{
			// 単語帳一覧取得
			listUrl.GET("/all", endpoint.GetBookListController)
			// 個別単語帳取得
			listUrl.GET("/individual", endpoint.GetIndividualBookController)
		}
		//認証API
		authUrl := ORIGIN_PATH.Group("/auth")
		{
			// ユーザー登録
			authUrl.POST("/register", endpoint.RegisterUserController)
			// 仮登録
			authUrl.POST("/temp", endpoint.RegisterTempUserController)
			// 仮登録完了
			authUrl.GET("/temp", endpoint.ValidateTempUserController)
			// ログイン
			authUrl.POST("/login", AuthMiddleware.LoginHandler)
			// トークンリフレッシュ
			authUrl.POST("/refresh", AuthMiddleware.RefreshHandler)
		}
	}

	router.NoRoute(AuthMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Println("claims: ", claims)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "PAGE NOT FOUND",
		})
	})

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err.Error())
	}
}
