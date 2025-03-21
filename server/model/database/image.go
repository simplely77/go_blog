package database

import (
	"gorm.io/gorm"
	"server/model/appTypes"
)

// Image 图片表
type Image struct {
	gorm.Model
	Name     string            `json:"name"`                       // 名称
	URL      string            `json:"url" gorm:"size:255;unique"` // 路径
	Category appTypes.Category `json:"category"`                   // 类别
	Storage  appTypes.Storage  `json:"storage"`                    // 存储类型
}
