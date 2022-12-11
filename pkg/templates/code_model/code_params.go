package code_model

const (
	CODING_RAW    = "RAW"
	CODING_HEX    = "Hex"
	CODING_Base64 = "Base64"
	CODING_GBK    = "GBK2UTF8"
)

type CodeParams struct {
	CodeType     string       // 编码分类
	CodeName     string       // 编码名称
	CodeStatus   bool         // 该编码是否开启
	CodeMode     string       // Encode/Decode
	CodeOptions  []CodeOption // 编码可选项切片
	InputCoding  string       // 输入编码
	OutputCoding string       // 输出编码
}

type CodeOption struct {
	KeyName string // 选项名
	Value   string // 值
	Coding  string // 值对应的编码
}
