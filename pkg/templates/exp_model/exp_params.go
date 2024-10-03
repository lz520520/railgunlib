package exp_model

// -------------------------------------exp响应信息结构体-------------------------------------
type ExpResult struct {
	Status bool   `json:"status"` // 利用是否成功
	Result string `json:"result"` // 执行返回结果

	RawResult string `json:"raw_result"` // http响应完整内容
	Err       string `json:"err"`        // 错误信息
}

// -------------------------------------exp发送参数集----------------------------------------------------
type ExpSendParams struct {
	BaseParam ExpBaseParam `json:"base_param"`
	Settings  ExpSettings  `json:"settings"`
}

type ExpBaseParam struct {
	App     string `json:"app"`
	Name    string `json:"name"`
	Payload string `json:"payload"`

	Target string `json:"-"` // 目标URL/IP等等
}
type ExpSettings struct {
	Timeout uint `json:"timeout"` // 超时时间，默认10秒
	Chunked bool `json:"chunked"` // 是否chunked编码发送

	ProxyEnable  bool   `json:"proxy_enable"`  // 是否开启代理
	ProxyAddress string `json:"proxy_address"` // 代理
	HttpHeader   string `json:"http_header"`   // 请求头部
	Charset      string `json:"charset"`       // 编写
	Debug        bool   `json:"debug"`         // 是否开启debug调试
}
