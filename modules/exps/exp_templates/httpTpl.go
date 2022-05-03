package exp_templates

import (
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

// -----------------------HTTP请求----------------------------------
func (self *ExpTemplate) HttpGet(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpGetWithoutRedirect(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

// -----------------------HTTP请求----------------------------------
func (self *ExpTemplate) HttpDelete(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpDeleteWithoutRedirect(target string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpGetWithSocket(target string, headers lzhttp.Header) (httpresp lzhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPostWithSocket(target string, data string, headers lzhttp.Header) (httpresp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPost(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPostWithoutRedirect(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPut(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPutWithoutRedirect(target, data string, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPostMulti(target string, postMultiParts []lzhttp.PostMultiPart, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPostMultiWithoutRedirect(target string, postMultiParts []lzhttp.PostMultiPart, headers lzhttp.Header) (resp lzhttp.HttpResp) {
	return
}
