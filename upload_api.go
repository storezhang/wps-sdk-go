package wps

import (
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	log `github.com/sirupsen/logrus`
)

func (w *Wps) UploadNetworkFile(url string) (rsp UploadFileRsp, err error) {
	var wpsRsp *resty.Response

	wpsRsp, err = NewResty().
		SetBody(UploadNetworkFileReq{Url: url}).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/httpFile", w.convertUrl()))

	if nil != err {
		log.WithFields(log.Fields{
			"wps":   w,
			"url":   url,
			"error": err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		log.WithFields(log.Fields{
			"wps":        w,
			"url":        url,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("上传文件失败")
	}

	return
}
