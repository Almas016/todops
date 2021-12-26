package models

import "time"

type Todo struct {
	Id          int       `gorm:"primaryKey"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
