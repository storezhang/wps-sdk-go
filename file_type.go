package wps

import (
	`path/filepath`
	`strings`
)

type (
	// SupportType 文件预览支持格式
	SupportType string
)

const (
	W SupportType = "w"
	S SupportType = "s"
	P SupportType = "p"
	// F SupportType = "f"
)

var wpsSupportExt = map[string]SupportType{
	// 文字文件
	"doc":  W,
	"dot":  W,
	"wps":  W,
	"wpt":  W,
	"docx": W,
	"dotx": W,
	"docm": W,
	"dotm": W,

	// 表格文件
	"xls":  S,
	"xlt":  S,
	"et":   S,
	"xlsx": S,
	"xltx": S,
	"xlsm": S,
	"xltm": S,

	// 演示文件
	"ppt":  P,
	"pptx": P,
	"pptm": P,
	"ppsx": P,
	"ppsm": P,
	"pps":  P,
	"potx": P,
	"potm": P,
	"dpt":  P,
	"dps":  P,
}

func CheckSupport(fileName string) (ok bool, st SupportType) {
	ext := filepath.Ext(fileName)
	ext = strings.ToLower(ext[1:])
	st, ok = wpsSupportExt[ext]

	return
}
