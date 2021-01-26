package wps

import (
	`strconv`
)

type (
	// BaseRsp 响应基类
	BaseRsp struct {
		// 状态码
		Code string `json:"code"`
		// 详细描述
		Msg string `json:"msg"`
		// 服务器时间
		ServerTime string `json:"servertime"`
	}
)

func (br *BaseRsp) ErrorCode() (code int) {
	var err error

	if code, err = strconv.Atoi(br.Code); nil != err {
		code = 1000
	}

	return
}
