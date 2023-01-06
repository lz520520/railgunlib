package poc_model

import (
	"net/http"
	"net/url"
	"time"
)

type PocProto string

var (
	PocHTTP PocProto = "HTTP"
	PocSMB  PocProto = "SMB"
	PocRDP  PocProto = "RDP"
)

type PocSendParamS struct {
	Target  string
	Port    string
	Cookie  string
	Timeout time.Duration
	Proxy   func(*http.Request) (*url.URL, error)
	Charset string
	Count   int
	Chunked bool
	PocName string
}

type PocPerPayloadResult struct {
	Status   bool   // 漏洞是否存在
	Messages string // 漏洞信息
	Err      error  // 错误
}

// poc 信息栏
type PocMsg struct {
	Name      string // 漏洞名称，如果为空，则直接用结构体名
	Author    string // 作者

	VulDate    string // 漏洞披露日期
	CreateDate string // 漏洞编写日期
	UpdateDate string // 漏洞修改日期
	//Deprecated
	Time string // 编写时间

	Range     string // 影响范围
	ID        string // 漏洞编号
	Describe  string // 漏洞描述
	Reference string // 引用
	VulType   string // 漏洞类型
}

