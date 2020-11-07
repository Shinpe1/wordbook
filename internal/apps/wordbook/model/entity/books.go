package entity

import (
	"time"
)

/** 単語帳マスタ */
type Book struct {
	ID        int64 // 単語帳ID
	UserID    int   // ユーザーID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time

	Contents []Content // 紐づくコンテンツ
}
