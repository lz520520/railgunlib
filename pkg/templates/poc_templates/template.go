package poc_templates

import (
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

type PocTemplate struct {
	Params poc_model.PocSendParamS
}

// 只有当前URL没有路径/目录时，会添加URI
func (self *PocTemplate) AddUri(target, uri string) (result string) {
	return
}

// 追加URL，基于当前目录
func (self *PocTemplate) AppendUri(target, uri string) (result string) {
	return
}
func (self *PocTemplate) GetHostname(target string) (hostname string) {

	return
}

func (self *PocTemplate) GetInitPocHeaders() (headers lzhttp.Header) {
	return
}
