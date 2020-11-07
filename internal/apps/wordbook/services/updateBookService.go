package services

import (
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
	"github.com/gin-gonic/gin"
)

func UpdateBookService(model *UpdateBookComp) {
	// err := inputValidaiton(model)
	// if err != nil {
	// 	panic(err.Error())
	// }

	db, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
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

	// books, contentsテーブルをそれぞれ更新します
	db.Model(&book).Update(&book)
}

/** テーブルに保存します */
func updateEntity(model *UpdateBookComp, ctx *gin.Context) {

}
