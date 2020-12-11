package main

import (
	"time"

	application "github.com/Shinpe1/wordbook_web/internal/apps/wordbook"
)

const TIMEZONE string = "Asia/Tokyo"

func main() {
	application.Run()
}

// タイムゾーンの指定
func init() {
	loc, err := time.LoadLocation(TIMEZONE)
	if err != nil {
		loc = time.FixedZone(TIMEZONE, 9*60*60)
	}
	time.Local = loc
}
