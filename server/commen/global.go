package commen

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	GVA_DB  *gorm.DB
	GVA_LOG *zap.Logger
)
