package room

import "time"

type Room struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	UserID          string    `json:"user_id" gorm:"not null;index"`
	RoomName        string    `json:"room_name" gorm:"not null"`
	RoomDescription *string   `json:"room_description" gorm:"type:text"`
	RoomTags        []string  `json:"room_tags" gorm:"type:text"`
	RoomTypeID      int       `json:"room_type_id" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	LastUpdateAt    time.Time `json:"last_update_at" gorm:"autoUpdateTime"`
	IsEnded         bool      `json:"is_ended" gorm:"not null;default:false"`
}

type RoomType struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
