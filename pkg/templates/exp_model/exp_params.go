package exp_model

import (
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
	"net/http"
	"net/url"
	"time"
)

// exp 信息栏
type ExpMsg struct {
	Author    string // 作者信息
	Time      string // 编写时间
	Range     string // 影响范围
	ID        string // CVE等编号
	Describe  string // 漏洞描述
	Details   string // 漏洞利用详细说明，每个payload都有一些操作步骤，可在此详细说明。
	Payload   string // 简易payload，用于提示用
	Reference string // 引用
	VulType   string // 漏洞类型
}

// -------------------------------------exp发送参数集----------------------------------------------------
type ExpSendParamS struct {
	ExpBaseOptions
	ExpNormalOptions

	Options ExpExtendOptions
}
type ExpBaseOptions struct {
	Target    string // 目标URL/IP等等
	Cmd       string // 执行的命令，这个可忽略，仅用于UI上提示
	Cookie    string // URL Cookie
	YsoGadget string // 反序列化所需要的gadget，暂时还没开放
}
type ExpNormalOptions struct {
	Timeout    time.Duration // 超时时间，默认10秒
	Chunked    bool          // 是否chunked编码发送
	OptRequest string        // 这个可忽略，后续移除

	Proxy   func(*http.Request) (*url.URL, error) // 代理
	Headers lzhttp.Header                         // 请求头部
	Charset string                                // 编写
	Debug   bool                                  // 是否开启debug调试
}

// 扩展选项
type ExpExtendOptions struct {
	Gadget    []string  // 所需要展示在UI上可选的gadget，暂时未开放
	VulParams VulParams // 漏洞信息

	CmdSubOptions    map[string]string // cmd子选项
	UploadSubOptions map[string]string // 上传子选项
}

// 漏洞名称等信息
type VulParams struct {
	VulMode    string
	VulName    string
	VulPayload string
}

// -------------------------------------exp响应信息结构体-------------------------------------
// 用于GetMsg/Cmd/Rerverse
type ExpResult struct {
	Cmd       string // 执行的cmd
	Result    string // 执行返回结果
	RawResult string // http响应完整内容
	Err       error  // 错误信息
}

// 用于Upload
type ExpUploadResult struct {
	Status    bool   // 是否上传成功
	RespPath  string // 上传后shell路径
	Msg       string // 额外信息
	RawResult string // http响应完整内容
	Err       error  // 错误信息
}
