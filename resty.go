package wps

import (
	`crypto/tls`

	`github.com/go-resty/resty/v2`
)

// NewResty Resty客户端
func NewResty() *resty.Request {
	return resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		R().
		ForceContentType(DefaultWpsContentType)
}
