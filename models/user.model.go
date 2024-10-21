package models

import "time"

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Date     time.Time `json:"date"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
