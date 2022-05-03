package exp_model

type UploadModeType int

const (
	UploadModeNormal = iota
	//UploadModeSegment
	//UploadModeMemShell
)

type ExpSubOption struct {
	CmdContent    string             // cmd默认内容
	CmdSubOptions []ExpSubOptionItem // cmd子选项自绘

	UploadPath       string             // 上传路径默认值
	UploadContent    string             // 上传内容默认值
	UploadSubOptions []ExpSubOptionItem // 上传子选项自绘

	Gadgets []string // gadget选项

	UploadModes []UploadModeType // 上传类型，插件暂时不提供其他选项
}

type ExpSubOptionItem struct {
	StaticText   string      // 标签名称
	Key          string      // 提取选项值所需要的参数
	Value        interface{} // 只有两种类型，string和[]string，分别对应Edit和ComboBox
	DefaultWidth int32       // 默认选项框长度，如果为0，则根据Value长度自适应。
	Position     int         // 位置，只有0/2两个值，因为cmd选项太多，如果需要第二行，则该值设置成2，
}
