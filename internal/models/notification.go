package models

import "github.com/google/uuid"

type Notification struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey"`
	Priority   string    `json:"priority"`
	Status     string    `json:"status"`
	Expiration int       `json:"expiration"`
	Scheduled  int       `json:"scheduled"`
	Type	   string    `json:"type"`
	
	MessageId uuid.UUID `json:"message_id"`
	Message    Message   `json:"message" gorm:"foreignKey:MessageId"`

	RecipientId uuid.UUID `json:"recipient_id"`
	Recipient  User      `json:"recipient" gorm:"foreignKey:RecipientId"`

	SenderId uuid.UUID `json:"sender_id"`
	Sender   User      `json:"sender" gorm:"foreignKey:SenderId"`
}