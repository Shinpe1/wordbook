package services

import (
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

func RegisterNewBookService(model *NewBook) {
	// 入力チェック
	inputValidation(model)

	// データベースに接続します
	db, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 現在時刻を取得
	now := time.Now()
	// Bookインスタンス作成
	newBook := Book{}
	newBook.UserID = model.UserId
	newBook.Title = model.Title
	newBook.CreatedAt = now
	newBook.UpdatedAt = now
	// Contentsインスタンス作成
	contents := []Content{}
	for i, _ := range model.Contents {
		content := Content{}
		content.BookID = newBook.ID
		content.UserID = newBook.UserID
		content.Head = model.Contents[i].Head
		content.Tail = model.Contents[i].Tail
		content.CreatedAt = now
		content.UpdatedAt = now
		contents = append(contents, content)
	}
	newBook.Contents = contents

	db.Create(&newBook)
}

/** 入力の検証を行います */
func inputValidation(model *NewBook) {

	// デフォルトは「無題」
	if model.Title == "" {
		model.Title = "無題"
	}

	// TODO: headとtailが両方とも空欄だったら削除したいんだけど、どうすればいいのかな？
}
