package models

import "github.com/google/uuid"

type Notification struct {
	Id         uuid.UUID `json:"id"`
	Priority   string    `json:"priority"`
	Recipient  string    `json:"recipient_id"`
	Message    Message   `json:"message"`
	Status     string    `json:"status"`
	Expiration int       `json:"expiration"`
	Scheduled  int       `json:"scheduled"`
	Type	   string    `json:"type"`
}