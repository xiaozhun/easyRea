package utils

import (
	"easyReq/global"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join("log", "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName("latest_log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
