package core

import "GoLog/router"

// 启动主程序
func RunServer() {
	Router := router.SetRouters()
	Router.Run(":8080")
}
