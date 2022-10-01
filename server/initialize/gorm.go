package initialize

import (
	"GoLog/commen"
	"GoLog/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库,主要创建连接，并根据数据初始化表格的tables
func InitializeDb() *gorm.DB {
	m := commen.GVA_CONFIG.Db
	if m.DbName == "" {
		return nil
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               m.UserName + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config,
		DefaultStringSize: 190, // string类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		panic(err) // 创建连接失败
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Article{},
		model.UserLike{},
		model.Comments{},
	)
	if err != nil {
		commen.GVA_LOG.Error("注册表格失败！", zap.Error(err))
	}
	commen.GVA_LOG.Info("注册表格成功")
}
