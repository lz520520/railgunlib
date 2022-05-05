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
	Status   bool
	Messages string
	Err      error
}

// poc 信息栏
type PocMsg struct {
	Name     string
	Author   string
	Time     string
	Range    string
	ID       string
	Describe string
}
