package commen

import (
	"GoLog/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ćšć±ćé
var (
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
