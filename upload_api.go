package wps

import (
	`fmt`
	`net/http`
	`net/url`

	`github.com/go-resty/resty/v2`
	`github.com/google/go-querystring/query`
	log `github.com/sirupsen/logrus`
)

// UploadNetworkFile 上传网络文件
// 可以是Http和Https地址，且只支持这两种地址
func (w *Wps) UploadNetworkFile(fileUrl string) (rsp UploadFileRsp, err error) {
	var (
		params url.Values
		wpsRsp *resty.Response
	)

	req := UploadNetworkFileReq{Url: fileUrl}
	if params, err = query.Values(req); nil != err {
		return
	}
	if wpsRsp, err = NewResty().
		SetFormDataFromValues(params).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/httpFile", w.convertUrl())); nil != err {
		log.WithFields(log.Fields{
			"url":     w.ApiUrl,
			"fileUrl": fileUrl,
			"error":   err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		err = newFileConvertError(wpsRsp.StatusCode())

		log.WithFields(log.Fields{
			"url":        w.ApiUrl,
			"fileUrl":    fileUrl,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("上传文件失败")
	}

	return
}
