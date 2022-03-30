package core

import (
	"easyReq/global"
	"easyReq/initialize"
	"easyReq/model"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	sPort := 17070
	// 打开文件
	file, _ := os.Open("config.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	conf := model.SysConfig{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("取配置端口Error:", err)
		conf.Port = 0
	}
	if conf.Port > 0 {
		sPort = conf.Port
	}
	address := fmt.Sprintf(":%d", sPort)
	// 统一返回格式初始化
	global.GVA_RES = conf.InitRes
	global.GVA_PATH = conf.InitPath
	// 初始化路由
	Router := initialize.Routers()
	s := initServer(address, Router)
	// 启动程序
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	默认访问地址:http://127.0.0.1%s
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
