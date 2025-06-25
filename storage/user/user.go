package user

import "time"

type User struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null;unique"`
	Salt        []byte    `json:"salt"`
	Password    string    `json:"password"`
	IsVerified  bool      `json:"is_verified" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"autoCreateTime"`

	Bio      *string    `json:"bio" gorm:"type:text"`
	Nickname *string    `json:"nickname" gorm:"type:text"`
	Avatar   *string    `json:"avatar" gorm:"type:text"`
	Location *string    `json:"location" gorm:"type:text"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Gender   string     `json:"gender" gorm:"not null;default:not-specified"`
	Role     string     `json:"role" gorm:"not null;default:user"`

	IsBanned bool `json:"is_banned" gorm:"not null;default:false"`
}
