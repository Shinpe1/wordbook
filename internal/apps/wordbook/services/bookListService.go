package services

import (
	"errors"
	"log"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	"github.com/Shinpe1/wordbook_web/settings/db"
)

/** 単語帳一覧を返します */
func GetListService(userId int) ([]Book, error) {
	log.Println("#GetListService start;")

	// データベースに接続します
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ユーザーIDが一致する単語帳一覧をすべて返します
	booksModel := []Book{}
	err = db.Find(&booksModel, "user_id=?", userId).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("DB access failed")
	}

	log.Println("#GetListService end;")
	return booksModel, nil
}

/** 個別の単語帳を返します */
func GetIndividualBookService(userId int, bookId int64) ([]Book, error) {
	log.Println("#GetIndividualBookService start;")
	// データベースに接続します
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	// 接続解除を遅延実行します
	defer db.Close()

	// ユーザーIDと単語帳IDが一致するコンテンツを返します
	var result []Book

	// contentsテーブルを最初に読み込んでおいてN+1問題を解消
	// 単語帳IDとユーザーIDが一致するデータをbooksから取得
	db.Preload("Contents").Find(&result, "id = ? AND user_id = ?", bookId, userId)
	// 取得したデータと紐づくcontentsテーブルのデータを取得
	for i, _ := range result {
		err = db.Model(result[i]).Related(&result[i].Contents).Error
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("DB access failed")
		}
	}

	log.Println("#GetIndividualBookService end;")

	return result, nil
}
