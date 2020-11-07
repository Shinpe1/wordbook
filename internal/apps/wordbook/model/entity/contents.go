package entity

import (
	"time"
)

/** コンテンツテーブル */
type Content struct {
	ID        int64
	BookID    int64
	UserID    int
	Head      string
	Tail      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
