package services

import (
	"errors"
	"log"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

func DeleteBookService(model *DeleteBookComp) error {
	log.Println("#DeleteBookService start;")
	log.Println("model : ", model)

	db, err := db.ConnectDB()
	if err != nil {
		return errors.New("couldn't connect database;")
	}

	tx := db.Begin()
	// 単語帳を丸ごと消去する
	if len(model.ContentsId) == 0 {
		err = tx.Where("user_id = ? AND id = ?", model.UserId, model.BookId).Delete(&Book{}).Error
	} else {
		// コンテンツを各個削除する
		for _, id := range model.ContentsId {
			err = tx.Delete(&Content{}, id).Error
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		tx.Rollback()
		return errors.New("couldn't delete records. Please try again")
	} else {
		tx.Commit()
	}

	log.Println("#DeleteBookService end;")
	return nil

}
