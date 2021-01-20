package wps

import (
	`encoding/json`
)

type (
	// ConvertReq 文件转换请求
	ConvertReq struct {
		// 上传后的文档Id
		Id string `json:"id"`
		// 目标文件格式，如果为空表示默认转换为默认的预览格式
		TargetFileFormat FormatType `json:"targetFileFormat"`
	}

	// ConvertRsp 文档转换响应
	ConvertRsp struct {
		BaseRsp

		// 数据
		Data struct {
			// Download 下载Id，如果不存在此字段表示文档不可下载
			Download string `json:"download"`
			// Preview 预览Id，如果不存在此字段表示该文档不可被预览
			Preview string `json:"preview"`
			// OfficeUrl Office预览Uri，仅在后台管理开启配置且为WebOffice预览时才存在
			OfficeUrl string `json:"officeUrl"`
		} `json:"data"`
	}
)

func (cr ConvertReq) String() string {
	jsonBytes, _ := json.MarshalIndent(cr, "", "    ")

	return string(jsonBytes)
}

func (fer ConvertRsp) String() string {
	jsonBytes, _ := json.MarshalIndent(fer, "", "    ")

	return string(jsonBytes)
}
