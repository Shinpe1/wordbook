package services

import (
	"errors"
	"log"

	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/entity"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/request"
	. "github.com/Shinpe1/wordbook_web/internal/apps/wordbook/model/response"
	"github.com/Shinpe1/wordbook_web/settings/db"
	"github.com/Shinpe1/wordbook_web/util"
)

func RegisterUserService(model RegisterUserRequest) error {
	log.Println("#RegisterUserService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	tx := db.Begin()
	var user User
	tx.Where("email = ?", model.Email).Find(&user)
	// ID = 0は初期値 = 存在しない
	if user.ID != 0 {
		tx.Rollback()
		return errors.New("Requested user is already exists")
	}

	user.Name = model.Name
	user.Email = model.Email
	encodedPassword := util.Encode(model.Password)
	user.Password = encodedPassword

	err = db.Select("Name", "Email", "Password").Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	log.Println("#RegisterUserService end;")
	return nil
}

func LoginUserService(model LoginUserRequest) (response LoginUserResponse, err error) {
	log.Println("#LoginUserService start;")

	db, err := db.ConnectDB()
	if err != nil {
		return LoginUserResponse{}, err
	}

	encodedPassword := util.Encode(model.Password)
	var user User
	db.Where("email = ? AND password = ?", model.Email, encodedPassword).Find(&user)
	if user.ID != 0 {
		// トークン生成
	} else {
		// ログイン失敗
		return LoginUserResponse{}, errors.New("Login Failed")
	}

	log.Println("#LoginUserService end;")
	return LoginUserResponse{}, nil
}
