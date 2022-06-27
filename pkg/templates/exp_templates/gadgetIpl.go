package exp_templates

import (
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
	"net/http"
	"time"
)

// 添加EncodeEcho系列Gadget请求头
func (self *ExpTemplate) AddEncodeCmdHeader(srcheaders lzhttp.Header, cmd string) {
}

// 添加Echo系列Gadget请求头
func (self *ExpTemplate) AddPlainCmdHeader(srcheaders lzhttp.Header, cmd string) {
}

// 检查是否回显利用成功
func (self *ExpTemplate) CheckRespHeader(headers http.Header) bool {
	return false
}

// 解析EncodeEcho系列响应数据
func (self *ExpTemplate) ParserEncodeCmdResult(rawResult string) (parserResult string, err error) {
	return parserResult, nil
}

// 解析Echo系列响应数据
func (self *ExpTemplate) ParserPlainCmdResult(rawResult string) (parserResult string, err error) {
	return parserResult, nil
}

// 回调延迟利用函数，判断是否存在漏洞
func (self *ExpTemplate) CheckGagdetWithSleep(fun func() error, delay time.Duration) (status bool, err error) {
	return
}
