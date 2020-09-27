package wps

import (
	log `github.com/sirupsen/logrus`
)

// ConvertByNetworkFile 转换网络文件
func (w *Wps) ConvertByNetworkFile(
	url string,
	targetFileFormat string,
) (previewId string, downloadId string, fileId string, err error) {
	// 上传文件
	var uploadRsp UploadFileRsp
	if uploadRsp, err = w.UploadNetworkFile(url); nil != err {
		return
	}
	if StatusOk != uploadRsp.Code {
		log.WithFields(log.Fields{
			"wps":   w,
			"rsp":   uploadRsp,
			"error": err,
		}).Error("上传文件出错")

		return
	}
	fileId = uploadRsp.Data.Id

	// 拿到预览结果
	var convertRsp ConvertRsp
	if convertRsp, err = w.Convert(fileId, targetFileFormat); nil != err {
		return
	}
	if StatusOk != convertRsp.Code {
		log.WithFields(log.Fields{
			"wps":   w,
			"rsp":   convertRsp,
			"error": err,
		}).Error("转换文件出错")

		return
	}
	previewId = convertRsp.Data.Preview
	downloadId = convertRsp.Data.Download

	return
}
