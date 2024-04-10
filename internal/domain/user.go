package domain

import "time"

type User struct {
	Uuid      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
