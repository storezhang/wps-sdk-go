package wps

import (
	`fmt`
	`net/http`

	log `github.com/sirupsen/logrus`
)

// Convert 文档转换
func (w *Wps) Convert(id string, formatType FormatType) (rsp ConvertRsp, err error) {
	var wpsRsp *resty.Response

	wpsRsp, err = NewResty().
		SetBody(ConvertReq{Id: id, TargetFileFormat: formatType}).
		SetResult(&rsp).
		Post(fmt.Sprintf("%s/api/convert", w.convertUrl()))

	if nil != err {
		log.WithFields(log.Fields{
			"wps":        w,
			"id":         id,
			"formatType": formatType,
			"error":      err,
		}).Error("转换文件出错")

		return
	}

	if http.StatusOK != wpsRsp.StatusCode() {
		log.WithFields(log.Fields{
			"wps":        w,
			"id":         id,
			"formatType": formatType,
			"statusCode": wpsRsp.StatusCode(),
		}).Error("转换文件失败")
	}

	return
}
