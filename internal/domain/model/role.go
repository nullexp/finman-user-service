package model

import "time"

type Role struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Permissions []string  `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
