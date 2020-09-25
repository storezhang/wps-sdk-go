package wps

import (
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	log `github.com/sirupsen/logrus`
)

func (w *Wps) Exist(id string, fileType FileType) (rsp FileExistRsp, err error) {
	var wpsRsp *resty.Response

	wpsRsp, err = NewResty().
		SetBody(FileExistReq{Id: id, Type: fileType}).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/verify/exists/file", w.convertUrl()))

	if nil != err {
		log.WithFields(log.Fields{
			"wps":   w,
			"id":    id,
			"type":  fileType,
			"error": err,
		}).Error("上传文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		log.WithFields(log.Fields{
			"wps":        w,
			"id":         id,
			"type":       fileType,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("上传文件失败")
	}

	return
}
