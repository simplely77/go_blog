package database

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Feedback 反馈表
type Feedback struct {
	gorm.Model
	UserUUID uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`               // 用户 uuid
	User     User      `json:"-" gorm:"foreignKey:UserUUID;references:UUID"` // 关联的用户
	Content  string    `json:"content"`                                      // 内容
	Reply    string    `json:"reply"`                                        // 回复
}
