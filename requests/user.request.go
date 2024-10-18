package requests

import "time"

type UserRequest struct {
	//ID       string    `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Address  string    `json:"address" binding:"required"`
	BornDate time.Time `json:"date" binding:"required"`
}
