package models

import "github.com/google/uuid"

type Message struct {
	Id    uuid.UUID `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Body  string `json:"body"`
}