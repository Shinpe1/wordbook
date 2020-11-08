package services

import (
	"log"
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
	"github.com/gin-gonic/gin"
)

func UpdateBookService(model *UpdateBookComp) error {
	log.Println("#UpdateBookService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	defer db.Close()

	// 現在時刻を取得
	now := time.Now()

	book := Book{}
	book.ID = model.BookId
	book.UserID = model.UserId
	book.Title = model.Title
	book.UpdatedAt = now
	contents := []Content{}
	for i, _ := range model.Contents {
		content := Content{}
		content.ID = model.Contents[i].ContentsId
		content.BookID = model.BookId
		content.UserID = model.UserId
		content.Head = model.Contents[i].Head
		content.Tail = model.Contents[i].Tail
		content.UpdatedAt = now
		contents = append(contents, content)
	}
	book.Contents = contents

	tx := db.Begin()
	// books, contentsテーブルをそれぞれ更新します
	err = tx.Model(&book).Update(&book).Error
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	log.Println("#UpdateBookService end;")

	return nil
}

/** テーブルに保存します */
func updateEntity(model *UpdateBookComp, ctx *gin.Context) {

}
