package services

import (
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

func RegisterNewBookService(request *NewBook) {
	// 入力チェック
	inputValidation(request)

	// データベースに接続します
	db, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	// 現在時刻を取得
	now := time.Now()
	// Bookインスタンス作成
	newBook := Book{}
	newBook.UserID = request.UserId
	newBook.Title = request.Title
	newBook.CreatedAt = now
	newBook.UpdatedAt = now
	// Contentsインスタンス作成
	contents := []Content{}
	for i, _ := range request.Contents {
		content := Content{}
		content.BookID = newBook.ID
		content.UserID = newBook.UserID
		content.Head = request.Contents[i].Head
		content.Tail = request.Contents[i].Tail
		content.CreatedAt = now
		content.UpdatedAt = now
		contents = append(contents, content)
	}
	newBook.Contents = contents

	// fmt.Println(contents)

	db.Create(&newBook)
	// db.Create(&contents)
}

/** 入力の検証を行います */
func inputValidation(request *NewBook) {

	// デフォルトは「無題」
	if request.Title == "" {
		request.Title = "無題"
	}

	// TODO: headとtailが両方とも空欄だったら削除したいんだけど、どうすればいいのかな？
}
