package services

import (
	"errors"
	"log"
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

func RegisterNewBookService(model *NewBook) error {
	log.Println("#RegisterNewBookService start;")
	// 入力チェック
	inputValidation(model)

	// データベースに接続します
	db, err := db.ConnectDB()
	if err != nil {
		return err
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

	tx := db.Begin()

	err = tx.Create(&newBook).Error
	if err != nil {
		log.Println("Couldn't create new book")
		tx.Rollback()
		return errors.New("Couldn't create new book")
	} else {
		tx.Commit()
	}

	log.Println("#RegisterNewBookService end;")

	return nil
}

/** 入力の検証を行います */
func inputValidation(model *NewBook) {

	// デフォルトは「無題」
	if model.Title == "" {
		model.Title = "無題"
	}

	// TODO: headとtailが両方とも空欄だったら削除したいんだけど、どうすればいいのかな？
}
