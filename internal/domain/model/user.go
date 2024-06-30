package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	RoleId    string    `json:"role_id"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
