package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `json:"id"`
	MongoDBId    string    `json:"mongodb_id"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Preferences  Preference `json:"preferences"`
	Active       bool      `json:"active"`
}