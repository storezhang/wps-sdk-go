package wps

import (
	"path/filepath"
	"strings"
)

type (
	// SupportType 文件预览支持格式
	SupportType string
	// FormatType 扩展类型
	FormatType string
)

const (
	SupportTypeWriter       SupportType = "w"
	SupportTypeSpreadsheets SupportType = "s"
	SupportTypePresentation SupportType = "p"
	SupportTypePdf          SupportType = "f"
)

const (
	FormatTypeNone FormatType = ""
	// 文字文件
	FormatTypeDoc  FormatType = "doc"
	FormatTypeDot  FormatType = "dot"
	FormatTypeWps  FormatType = "wps"
	FormatTypeWpt  FormatType = "wpt"
	FormatTypeDocx FormatType = "docx"
	FormatTypeDotx FormatType = "dotx"
	FormatTypeDocm FormatType = "docm"
	FormatTypeDotm FormatType = "dotm"

	// 表格文件
	FormatTypeXls  FormatType = "xls"
	FormatTypeXlt  FormatType = "xlt"
	FormatTypeEt   FormatType = "et"
	FormatTypeXlsx FormatType = "xlsx"
	FormatTypeXltx FormatType = "xltx"
	FormatTypeXlsm FormatType = "xlsm"
	FormatTypeXltm FormatType = "xltm"

	// 演示文件
	FormatTypePpt  FormatType = "ppt"
	FormatTypePptx FormatType = "pptx"
	FormatTypePptm FormatType = "pptm"
	FormatTypePpsx FormatType = "ppsx"
	FormatTypePpsm FormatType = "ppsm"
	FormatTypePps  FormatType = "pps"
	FormatTypePotx FormatType = "potx"
	FormatTypePotm FormatType = "potm"
	FormatTypeDpt  FormatType = "dpt"
	FormatTypeDps  FormatType = "dps"

	// Pdf文件
	FormatTypePdf FormatType = "pdf"
	FormatTypeOfd FormatType = "ofd"
)

var wpsSupportExt = map[FormatType]SupportType{
	// 文字文件
	FormatTypeDoc:  SupportTypeWriter,
	FormatTypeDot:  SupportTypeWriter,
	FormatTypeWps:  SupportTypeWriter,
	FormatTypeWpt:  SupportTypeWriter,
	FormatTypeDocx: SupportTypeWriter,
	FormatTypeDotx: SupportTypeWriter,
	FormatTypeDocm: SupportTypeWriter,
	FormatTypeDotm: SupportTypeWriter,

	// 表格文件
	FormatTypeXls:  SupportTypeSpreadsheets,
	FormatTypeXlt:  SupportTypeSpreadsheets,
	FormatTypeEt:   SupportTypeSpreadsheets,
	FormatTypeXlsx: SupportTypeSpreadsheets,
	FormatTypeXltx: SupportTypeSpreadsheets,
	FormatTypeXlsm: SupportTypeSpreadsheets,
	FormatTypeXltm: SupportTypeSpreadsheets,

	// 演示文件
	FormatTypePpt:  SupportTypePresentation,
	FormatTypePptx: SupportTypePresentation,
	FormatTypePptm: SupportTypePresentation,
	FormatTypePpsx: SupportTypePresentation,
	FormatTypePpsm: SupportTypePresentation,
	FormatTypePps:  SupportTypePresentation,
	FormatTypePotx: SupportTypePresentation,
	FormatTypePotm: SupportTypePresentation,
	FormatTypeDpt:  SupportTypePresentation,
	FormatTypeDps:  SupportTypePresentation,

	// Pdf文件
	FormatTypePdf: SupportTypePdf,
	FormatTypeOfd: SupportTypePdf,
}

func CheckSupport(filename string) (ok bool, st SupportType) {
	if "" == filename {
		return
	}
	ext := filepath.Ext(filename)
	if "" == ext {
		return
	}
	ext = strings.ToLower(ext[1:])
	st, ok = wpsSupportExt[FormatType(ext)]

	return
}
