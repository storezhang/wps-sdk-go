package wps

import (
	`fmt`
	`net/url`
	`strings`

	`github.com/google/go-querystring/query`
)

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
