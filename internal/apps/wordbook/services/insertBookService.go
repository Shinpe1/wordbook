package services

import (
	"errors"
	"log"
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

func InsertBookService(model *InsertBookComp) error {
	log.Println("#InsertBookService start;")
	db, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	tx := db.Begin()

	log.Println("requested model: ", model)

	defer db.Close()

	// forで繰り返してレコードの挿入
	// 公式docにあるバルクインサートがなぜかできない...
	now := time.Now()
	for _, c := range model.Contents {
		content := Content{}
		content.UserID = model.UserId
		content.BookID = model.BookId
		content.Head = c.Head
		content.Tail = c.Tail
		content.CreatedAt = now
		content.UpdatedAt = now

		err = tx.Create(&content).Error
	}
	if err != nil {
		tx.Rollback()
		return errors.New("couldn't save records")
	}
	tx.Commit()

	log.Println("#InsertBookService end;")
	return nil
}
