package lzhttp

import (
	"net/http"
)

type HttpResp struct {
	Resp        *http.Response
	Header      string
	Feature     string
	Title       string
	Body        string
	RawFullResp string
	Err         error
}

type PostMultiPart struct {
	FieldName   string
	FileName    string
	ContentType string
	Content     []byte
}
