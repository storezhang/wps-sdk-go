package wps

import (
	`encoding/json`
)

const (
	// 上传文件
	FileTypeUpload FileType = "UPLOAD"
	// 下载文件
	FileTypeDownload FileType = "DOWNLOAD"
	// 预览文件
	FileTypePreview FileType = "PREVIEW"
)

type (
	// FileType 文件类型
	FileType string
	// FileExistReq 文档是否存在的请求
	FileExistReq struct {
		// 文件Id
		Id string `json:"id"`
		// 文件类型
		Type FileType `json:"type"`
	}

	FileExistRsp struct {
		BaseRsp

		// 数据
		Data struct {
			// 文件是否存在
			ExistsFile bool `json:"existsFile"`
		} `json:"data"`
	}
)

func (fer FileExistReq) String() string {
	jsonBytes, _ := json.MarshalIndent(fer, "", "    ")

	return string(jsonBytes)
}

func (fer FileExistRsp) String() string {
	jsonBytes, _ := json.MarshalIndent(fer, "", "    ")

	return string(jsonBytes)
}
