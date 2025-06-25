package chat

import "time"

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"not null;index"`
	RoomID    string    `json:"room_id" gorm:"not null;index"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
