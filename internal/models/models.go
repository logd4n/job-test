package models

import "time"

type ResponseChat struct {
	ID       uint
	Messages []Message
}

type Chat struct {
	ID         uint      `gorm:"primaryKey"`
	Title      string    `gorm:"size:200; not_null"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Messages   []Message `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"`
}

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	Chat_ID    uint      `gorm:"not_null;index"`
	Text       string    `gorm:"size:5000;not_null"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Chat       Chat      `gorm:"foeignKey:Chat_ID;constraint:OnDelete:CASCADE;"`
}
