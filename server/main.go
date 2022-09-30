package main

import (
	"GoLog/commen"
	"GoLog/core"
	"GoLog/initialize"
)

func main() {
	// 首先读取配置,之后初始化全局变量,最后开启服务
	commen.GVA_VP = core.Viper() // 全局配置读取到commen.GVA_CONFIG,并实时监控文件变化
	commen.GVA_LOG = core.Zap()
	// 初始化数据库
	commen.GVA_DB = initialize.InitializeDb()

	if commen.GVA_DB != nil {
		commen.GVA_LOG.Info("连接数据库成功")
		// 初始化表格
		initialize.MigrateTables(commen.GVA_DB)
		db, _ := commen.GVA_DB.DB()
		defer db.Close()
	}
	// 开启服务
	core.RunServer()

}
