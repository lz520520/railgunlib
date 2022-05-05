package poc_templates

import (
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

// -----------------------HTTP请求----------------------------------
func (self *PocTemplate) HttpGet(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpGetWithoutRedirect(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpDelete(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpDeleteWithoutRedirect(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpGetWithSocket(target string, headers lzhttp.Header) (httpresp lzhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostWithSocket(target string, data string, headers lzhttp.Header) (httpresp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPost(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPostWithoutRedirect(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPut(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPutWithoutRedirect(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostMulti(target string, postMultiParts []lzhttp.PostMultiPart, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostMultiWithoutRedirect(target string, postMultiParts []lzhttp.PostMultiPart, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
