package wps

import (
	`fmt`
	`net/url`
	`strings`

	`github.com/google/go-querystring/query`
)

// PreviewUrl 获得预览地址
func (w *Wps) PreviewUrl(reader Reader) (previewUrl string, err error) {
	var values url.Values

	previewUrl = fmt.Sprintf("%s/reader", w.previewUrl())
	if values, err = query.Values(reader); nil != err {
		return
	}
	params := values.Encode()
	if "" != strings.TrimSpace(params) {
		previewUrl = fmt.Sprintf("%s?%s", previewUrl, params)
	}

	return
}

// InnerPreviewUrl 获得预览地址，去掉外层iframe后的地址
func (w *Wps) InnerPreviewUrl(filename string, reader Reader) (previewUrl string, err error) {
	var values url.Values

	if values, err = query.Values(reader); nil != err {
		return
	}
	params := values.Encode()

	if ok, st := CheckSupport(filename); ok {
		previewUrl = fmt.Sprintf("%s/%s/%s?%s", w.innerPreviewUrl(), st, reader.File, params)
	}

	return
}

// DownloadUrl 获得下载地址
func (w *Wps) DownloadUrl(id string) (downloadUrl string, err error) {
	downloadUrl = fmt.Sprintf("%s/api/download/%s", w.convertUrl(), id)

	return
}
