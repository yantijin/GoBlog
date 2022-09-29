package commen

import (
	"GoLog/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
