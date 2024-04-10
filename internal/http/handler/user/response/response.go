package response

import "time"

type Register struct {
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Token string `json:"token"`
}
