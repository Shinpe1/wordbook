package db

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB接続設定
func ConnectDB() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	// PROTOCOL := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")
	// DBNAME := "wordbook"

	// datetime型をtime.Timeで受け取れるようにするため,parseTime=trueを指定している
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&reconnect=true&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return db, errors.New("Can't establish DB connection;")
	}

	// 発行したSQLを標準出力に出す
	db.LogMode(true)

	return db, err
}
