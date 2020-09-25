package wps

import (
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	log `github.com/sirupsen/logrus`
)

func (w *Wps) Convert(id string, targetFileFormat string) (rsp ConvertRsp, err error) {
	var wpsRsp *resty.Response

	wpsRsp, err = NewResty().
		SetBody(ConvertReq{Id: id, TargetFileFormat: targetFileFormat}).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/convert", w.convertUrl()))

	if nil != err {
		log.WithFields(log.Fields{
			"wps":              w,
			"id":               id,
			"targetFileFormat": targetFileFormat,
			"error":            err,
		}).Error("转换文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		log.WithFields(log.Fields{
			"wps":              w,
			"id":               id,
			"targetFileFormat": targetFileFormat,
			"statusCode":       wpsRsp.StatusCode(),
		}).Error("转换文件失败")
	}

	return
}
