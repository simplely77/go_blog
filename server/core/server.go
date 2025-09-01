package core

import (
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"server/service"
)

type server interface {
	ListenAndServe() error
}

// RunServer 用于启动服务器
func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()

	service.LoadAll()

	// 初始化服务器并启动
	s := initServer(addr, Router)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
