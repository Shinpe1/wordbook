package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	Name      sql.NullString
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
