package main

import (
	"easyReq/core"
	"easyReq/global"
)

func main() {
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	core.RunWindowsServer()
}
