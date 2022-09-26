package commen

import (
	"time"

	"gorm.io/gorm"
)

// 公用数据结构

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
