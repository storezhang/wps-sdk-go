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
	Writer       SupportType = "w"
	Spreadsheets SupportType = "s"
	Presentation SupportType = "p"
	// F SupportType = "f"
)

var wpsSupportExt = map[string]SupportType{
	// 文字文件
	"doc":  Writer,
	"dot":  Writer,
	"wps":  Writer,
	"wpt":  Writer,
	"docx": Writer,
	"dotx": Writer,
	"docm": Writer,
	"dotm": Writer,

	// 表格文件
	"xls":  Spreadsheets,
	"xlt":  Spreadsheets,
	"et":   Spreadsheets,
	"xlsx": Spreadsheets,
	"xltx": Spreadsheets,
	"xlsm": Spreadsheets,
	"xltm": Spreadsheets,

	// 演示文件
	"ppt":  Presentation,
	"pptx": Presentation,
	"pptm": Presentation,
	"ppsx": Presentation,
	"ppsm": Presentation,
	"pps":  Presentation,
	"potx": Presentation,
	"potm": Presentation,
	"dpt":  Presentation,
	"dps":  Presentation,
}

func CheckSupport(filename string) (ok bool, st SupportType) {
	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext[1:])
	st, ok = wpsSupportExt[ext]

	return
}
