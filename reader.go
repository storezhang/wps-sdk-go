package wps

// Reader 文档预览
type Reader struct {
	// 文件预览Id
	File string `json:"file" url:"file" validate:"required"`
	// 水印文字
	MarkText string `json:"markText" url:"markText"`
	// 水印颜色，十六进制形式，比如#DB7093
	// 默认值是#A6A6A6
	MarkColor string `json:"markColor" url:"markColor"`
	// 水印透明度，可取值是0到1，0为完全透明
	// 默认值是0.6
	MarkTransparent float32 `json:"markTransparent" url:"markTransparent"`
	// 字体CSS样式，比如bold 20px Serif
	MarkFontCssStyle string `json:"markFontCssStyle" url:"markFontCssStyle"`
	// 水印旋转度，可取值是1到360
	MarkRotate int `json:"markRotate" url:"markRotate"`
	// 水印水平距离，可取值是1到100
	MarkHorizontal int8 `json:"markHorizontal" url:"markHorizontal"`
	// 水印垂直距离，可取值是1到100
	MarkVertical int8 `json:"markVertical" url:"markVertical"`
	// 是否可打印，默认false
	IsPrint bool `json:"isPrint" url:"isPrint"`
	// 是否可以被下载，默认为false
	IsDownload bool `json:"isDownload" url:"isDownload"`
	// 页面内容是否可以被复制，默认false
	IsCopy bool `json:"isCopy" url:"isCopy"`
	// 从低到高位，每位的取值是0或者1，依次代表如下
	// 第0位：默认1，0-最终状态，1- 原始状态
	// 第1位：默认1（显示），是否显示标记
	// 第2位和第3位组合使用：00是批注框显示修订者，10是嵌入方式显示，01及11为批注框方式
	// 第4位：默认1（显示），是否显示评论
	// 第5位：默认1（显示），是否显示插入和删除
	// 第6位：默认1（显示），是否显示格式修订
	WpsPreview string `json:"wpsPreview" url:"wpsPreview"`
}
