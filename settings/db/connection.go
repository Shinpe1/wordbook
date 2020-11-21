package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB接続設定
func ConnectDB() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOSTNAME := os.Getenv("DB_HOSTNAME")
	DBNAME := os.Getenv("DB_NAME")

	// datetime型をtime.Timeで受け取れるようにするため,parseTime=trueを指定している
	CONNECT := USER + ":" + PASS + "@" + HOSTNAME + ")/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	log.Println("connect : " + CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return db, err
	}

	// 発行したSQLを標準出力に出す
	db.LogMode(true)

	return db, err
}
