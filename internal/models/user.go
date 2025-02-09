package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	MongoDBId    string    `json:"mongodb_id"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Preference   string    `json:"preference"`
	Active       bool      `json:"active"`
}