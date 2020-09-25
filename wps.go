package wps

import (
	`fmt`
)

const (
	// 默认的Content Type
	DefaultWpsContentType string = "application/x-www-form-urlencoded"

	// 状态码
	// 成功
	WpsStatusOk string = "200"
	// 参数错误
	WpsStatusParamsError string = "400"
)

// Wps 金山文档
type Wps struct {
	// 服务器地址
	Host string `json:"ip"`
	// 服务器端口
	Port int `json:"port"`
	// 协议，支持HTTP和HTTPS
	Scheme string `default:"http" json:"scheme"`
	// 文档预览前缀
	PreviewPrefix string `default:"web" json:"previewPrefix"`
	// 文档转换前缀
	ConvertPrefix string `default:"web-preview" json:"convertPrefix"`
}

func (w *Wps) previewUrl() string {
	return fmt.Sprintf("%s://%s:%d/%s", w.Scheme, w.Host, w.Port, w.PreviewPrefix)
}

func (w *Wps) convertUrl() string {
	return fmt.Sprintf("%s://%s:%d/%s", w.Scheme, w.Host, w.Port, w.ConvertPrefix)
}
