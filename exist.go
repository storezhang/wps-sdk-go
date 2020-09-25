package wps

type (
	// FileExistReq 文档是否存在的请求
	FileExistReq struct {
		// 文件Id
		Id string `json:"id"`
		// 文件类型
		Type string `json:"type"`
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
