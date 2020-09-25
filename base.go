package wps

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
