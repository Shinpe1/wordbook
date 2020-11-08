package wordbook

import (
	"github.com/Shinpe1/wordbook_web/internal/apps/wordbook/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// コンテクストパス
	originPath := router.Group("/api/v1.0")
	{
		// 取得API
		listAPI := originPath.Group("/list")
		{
			// 単語帳一覧取得
			listAPI.GET("/all", controller.GetBookListController)
			// 個別単語帳取得
			listAPI.GET("/individual", controller.GetIndividualBookController)
		}
		// 登録API
		registerUrl := originPath.Group("/register")
		{
			// 単語帳新規登録
			registerUrl.POST("", controller.RegisterBookController)
		}
		// 更新API
		updateUrl := originPath.Group("/update")
		{
			// 単語帳更新
			updateUrl.POST("", controller.UpdateBookController)
		}
		// 追加API
		insertUrl := originPath.Group("/insert")
		{
			// 単語帳追加
			insertUrl.POST("", controller.InsertBookController)
		}
		// 削除API
		deleteUrl := originPath.Group("/delete")
		{
			// 単語帳削除
			deleteUrl.POST("", controller.DeleteBookController)
		}
		// 認証API
		// authAPI := rouer.Group("/auth")
		// {
		// 	/ ユーザー登録
		// 	authAPI.POST("/", controller.RegisterUserController)
		// 	// ログイン
		// 	authAPI.POST("login", controller.LoginUserController)
		// 	// トークンリフレッシュ
		// 	authAPI.POST("/refresh", controller.RefreshTokenController)
		// }
	}

	router.Run(":8000")
}
