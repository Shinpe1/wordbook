package entity

import "time"

type Temp struct {
	Email    string
	ExpAt    time.Time
	Token    string
	IsActive bool
}
