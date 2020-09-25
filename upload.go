package wps

import (
	`encoding/json`
)

type (
	// UploadByLocalFileReq 本地文件上传
	UploadLocalFileReq struct {
		// 文件
		File string `json:"file"`
	}

	// UploadNetworkFileReq HTTP/HTTPS网络文件上传
	UploadNetworkFileReq struct {
		// 网络文件地址
		Url string `json:"url" url:"url"`
	}

	// UploadFileRsp 文件上传返回结果
	UploadFileRsp struct {
		BaseRsp

		// 数据
		Data struct {
			// WPS文件编号
			Id string `json:"id"`
		} `json:"data"`
	}
)

func (ur UploadLocalFileReq) String() string {
	jsonBytes, _ := json.MarshalIndent(ur, "", "    ")

	return string(jsonBytes)
}

func (ur UploadNetworkFileReq) String() string {
	jsonBytes, _ := json.MarshalIndent(ur, "", "    ")

	return string(jsonBytes)
}

func (ur UploadFileRsp) String() string {
	jsonBytes, _ := json.MarshalIndent(ur, "", "    ")

	return string(jsonBytes)
}
