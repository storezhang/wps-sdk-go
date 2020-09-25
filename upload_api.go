package wps

import (
	`fmt`
	`net/http`
	`net/url`

	`github.com/go-resty/resty/v2`
	`github.com/google/go-querystring/query`
	log `github.com/sirupsen/logrus`
)

func (w *Wps) UploadNetworkFile(fileUrl string) (rsp UploadFileRsp, err error) {
	var (
		params url.Values
		wpsRsp *resty.Response
	)

	req := UploadNetworkFileReq{Url: fileUrl}
	if params, err = query.Values(req); nil != err {
		return
	}
	wpsRsp, err = NewResty().
		SetFormDataFromValues(params).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/httpFile", w.convertUrl()))

	if nil != err {
		log.WithFields(log.Fields{
			"wps":     w,
			"fileUrl": fileUrl,
			"error":   err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		log.WithFields(log.Fields{
			"wps":        w,
			"fileUrl":    fileUrl,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("上传文件失败")
	}

	return
}
