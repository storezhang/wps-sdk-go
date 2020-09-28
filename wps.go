package wps

import (
	`encoding/json`
	`fmt`
)

const (
	// 状态码
	// 成功
	StatusOk string = "200"
	// 参数错误
	// StatusParamsError string = "400"
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
	// 内部文档预览前缀
	InnerPreviewPrefix string `default:"office" json:"innerPreviewPrefix"`
	// 文档转换前缀
	ConvertPrefix string `default:"web-preview" json:"convertPrefix"`
}

func (w Wps) String() string {
	jsonBytes, _ := json.MarshalIndent(w, "", "    ")

	return string(jsonBytes)
}

func (w *Wps) previewUrl() string {
	return fmt.Sprintf("%s://%s:%d/%s", w.Scheme, w.Host, w.Port, w.PreviewPrefix)
}

func (w *Wps) innerPreviewUrl() string {
	return fmt.Sprintf("%s://%s:%d/%s", w.Scheme, w.Host, w.Port, w.InnerPreviewPrefix)
}

func (w *Wps) convertUrl() string {
	return fmt.Sprintf("%s://%s:%d/%s", w.Scheme, w.Host, w.Port, w.ConvertPrefix)
}
