package requests

import "time"

type UserRequest struct {
	//ID       string    `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Address  string    `json:"address" binding:"required"`
	BornDate time.Time `json:"date" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
