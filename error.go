package wps

import (
	`github.com/storezhang/gox`
)

func newFileConvertError(code int) *gox.CodeError {
	return &gox.CodeError{ErrorCode: gox.ErrorCode(code), Message: "转换文件失败"}
}

func newFileUploadError(code int) *gox.CodeError {
	return &gox.CodeError{ErrorCode: gox.ErrorCode(code), Message: "上件上传失败"}
}
