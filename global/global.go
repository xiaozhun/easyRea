package global

import (
	"go.uber.org/zap"
)

var (
	GVA_LOG *zap.Logger

	/* 全局测试开关，用于一些测试场景，避免真实的发送短信和语音 */
	TestMode bool
	GVA_RES  string
	GVA_PATH string
)
