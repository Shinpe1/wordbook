package auth

import (
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	"github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/token"
	"github.com/Shinpe1/wordbook_web/settings/db"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/** トークン生成・検証 */
var AuthMiddleware *jwt.GinJWTMiddleware

func init() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "Asia/Tokyo",
		Key:           []byte(os.Getenv("SECRET_KEY")),
		Timeout:       time.Minute * 30,
		MaxRefresh:    time.Hour,
		TimeFunc:      time.Now,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		// Authenticatorの返り値がこのメソッドに渡される
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*token.TokenClaims); ok {
				// tokenに含めるmap型のclaimsを返す
				return jwt.MapClaims{
					"UserId": v.UserId,
					"Iss":    v.Iss,
					"Iat":    v.Iat,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx *gin.Context) interface{} {
			// claimsから情報を抽出する
			claims := jwt.ExtractClaims(ctx)
			return &token.TokenClaims{
				UserId: int(claims["UserId"].(float64)), // なぜかclaimsの中でfloat64型になってる...
				Iss:    claims["Iss"].(string),
				Iat:    claims["Iat"].(string),
			}
		},
		// トークンからデコードしたUserIdとクエリに含まれるuserIdの検証
		Authorizator: func(data interface{}, ctx *gin.Context) bool {
			method := ctx.Request.Method
			var userId int
			switch method {
			case "POST":
				userId, err = strconv.Atoi(ctx.PostForm("userId"))
			case "GET":
				userId, err = strconv.Atoi(ctx.Query("userId"))
			}
			v, ok := data.(*token.TokenClaims)
			if ok && userId == v.UserId {
				return true
			}
			return false
		},

		// ログイン画面からPOSTされたデータの検証
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var user request.LoginUserRequest

			err := ctx.ShouldBindJSON(&user)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			email := user.Email
			password := user.Password

			db, err := db.ConnectDB()
			if err != nil {
				return nil, err
			}
			defer db.Close()
			isUser := User{}
			db.Find(&isUser, "email = ?", email)

			// ユーザーIDとパスワードでログイン
			if email == isUser.Email && password == isUser.Password {
				return &token.TokenClaims{
					UserId: isUser.ID,
					Iss:    "wordbook",
					Iat:    time.Now().String(),
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// 401エラー
		Unauthorized: func(ctx *gin.Context, code int, message string) {
			ctx.JSONP(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// ログイン時のレスポンス
		LoginResponse: func(ctx *gin.Context, code int, token string, expire time.Time) {
			ctx.JSONP(code, gin.H{
				"code":  code,
				"token": token,
			})
		},

		// トークンリフレッシュ時のレスポンス
		RefreshResponse: func(ctx *gin.Context, code int, token string, expire time.Time) {
			ctx.JSONP(code, gin.H{
				"code":  code,
				"token": token,
			})

		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := AuthMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
