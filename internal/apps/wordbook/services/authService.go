package services

import (
	"errors"
	"log"
	"os"
	"time"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	"github.com/Shinpe1/wordbook_web/settings/db"
	"github.com/Shinpe1/wordbook_web/util"
	"github.com/jinzhu/gorm"
)

// 有効期限 ( 60分後 )
const EXPIRATION time.Duration = 60

// 仮登録
func RegisterTempUserService(model TempUserRequest) error {
	log.Println("#RegisterTempUserService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	tx := db.Begin()
	var temp Temp

	temp.Email = model.Email
	// 現在時刻から1時間後
	temp.ExpAt = time.Now().Add(EXPIRATION * time.Minute)
	// パスワードと同じアルゴリズムでトークンを生成
	temp.Token = util.Encode(os.Getenv("URL_SECRET") + temp.Email)
	// まだ非アクティブ
	temp.IsActive = false

	err = util.SendMail(temp.Email, temp.Token)
	if err != nil {
		return err
	}

	// メールアドレスで検索し、すでにあればそれは削除する
	err = tx.Where("email = ?", temp.Email).Delete(&temp).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// レコードの挿入
	err = tx.Create(&temp).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	log.Println("#RegisterTempUserService end;")
	return nil

}

// 本登録
func RegisterUserService(model RegisterUserRequest) error {
	log.Println("#RegisterUserService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	tx := db.Begin()
	var temp Temp
	err = tx.Where("token = ? AND is_active = true", model.Token).Take(&temp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// レコードがない場合
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	var user User
	user.Name = model.Name
	user.Email = temp.Email
	encodedPassword := util.Encode(model.Password)
	user.Password = encodedPassword

	err = db.Select("Name", "Email", "Password").Create(&user).Error
	if err != nil {
		log.Println("Failed to insert new user record")
		log.Println(err.Error())
		tx.Rollback()
		return err
	}

	err = tx.Delete(&temp).Error
	if err != nil {
		log.Println("Failed to delete remaining temp record")
		log.Println(err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	log.Println("#RegisterUserService end;")
	return nil
}

func ValidateTempTokenService(token string) (bool, error) {
	log.Println("#ValidateTempUserService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return false, err
	}

	tx := db.Begin()
	var temp Temp
	err = tx.Where("token = ? AND is_active = false", token).Take(&temp).Error

	// 現在時刻
	now := time.Now()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// レコードがない場合
		tx.Rollback()
		return false, nil
	} else if temp.ExpAt.Before(now) {
		// expAtが現在時刻よりも後の時刻ならNG. レコード消去
		tx.Delete(&temp)
		tx.Commit()
		return false, nil
	}

	temp.IsActive = true
	// 仮登録レコードのis_activeをtrueに
	err = tx.Model(&temp).Where("token = ?", token).Update("is_active", true).Error
	if err != nil {
		log.Println("Failed to update user status")
		log.Println(err.Error())
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	log.Println("#ValidateTempUserService end;")
	return true, nil
}
