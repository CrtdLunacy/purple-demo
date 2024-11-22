package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone     string `gorm:"index"`
	SessionId string `gorm:"session_id"`
	Code      int    `gorm:"code"`
}
