package wps

import (
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	log `github.com/sirupsen/logrus`
)

// Exist 判断文件是否存在
func (w *Wps) Exist(id string, fileType FileType) (exist bool, err error) {
	var (
		wpsRsp *resty.Response
		rsp    FileExistRsp
	)

	if wpsRsp, err = NewResty().
		SetBody(FileExistReq{Id: id, Type: fileType}).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/verify/exists/file", w.convertUrl())); nil != err {
		log.WithFields(log.Fields{
			"url":   w.ApiUrl,
			"id":    id,
			"type":  fileType,
			"error": err,
		}).Error("判断文件是否存在出错")

		return
	}

	log.WithFields(log.Fields{
		"url":      w.ApiUrl,
		"id":       id,
		"type":     fileType,
		"response": wpsRsp.String(),
	}).Debug("判断文件是否存在返回结果")

	if http.StatusNotFound == wpsRsp.StatusCode() {
		exist = false
	} else {
		exist = rsp.Data.ExistsFile
	}

	return
}
