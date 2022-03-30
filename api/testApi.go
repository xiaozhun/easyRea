package api

import (
	"easyReq/global"
	"easyReq/model"
	"easyReq/model/request"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	result, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		global.GVA_LOG.Error("未获取到body数据")
	}
	global.GVA_LOG.Sugar().Info("req data:", string(result))

	if global.GVA_RES != "" {
		para := make(map[string]interface{})
		json.Unmarshal([]byte(global.GVA_RES), &para)
		global.GVA_LOG.Sugar().Info("print return:", global.GVA_RES)
		c.JSON(
			http.StatusOK,
			para,
		)
	} else {
		// 返回的格式化数据，和实际返回应保持一致
		restr, _ := json.Marshal(model.Response{
			0,
			"Success",
			"test is ok",
		})
		global.GVA_LOG.Sugar().Info("print return", restr)
		c.JSON(
			http.StatusOK,
			model.Response{
				0,
				"Success",
				"test is ok",
			},
		)
	}
}

// 处理短信回调
func NoteSmsRecord(c *gin.Context) {
	var req request.NoteSmsRecordReq
	c.ShouldBindJSON(&req)
	marshal, _ := json.Marshal(req)
	global.GVA_LOG.Info("监听收到短信记录===》 " + string(marshal))
	c.JSON(
		http.StatusOK,
		model.Response{
			0,
			"Success",
			"",
		},
	)
}
