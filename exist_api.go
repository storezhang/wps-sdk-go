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
		}).Error("判断文件是否存在出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		err = newFileExistError(wpsRsp.StatusCode())

		log.WithFields(log.Fields{
			"wps":        w,
			"id":         id,
			"type":       fileType,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("判断文件是否存在失败")

		return
	}

	if StatusOk != rsp.Code {
		err = newFileExistError(rsp.ErrorCode())

		return
	}
	exist = rsp.Data.ExistsFile

	return
}
