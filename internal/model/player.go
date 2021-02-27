package model

import "time"

type Player struct {
	ID        *string    `json:"id"`
	UserId    *string    `json:"user_id"`
	Name      *string    `json:"name"`
	Level     *string    `json:"level"`
	Job       *string    `json:"job"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
