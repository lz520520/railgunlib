package exp_templates

/*
ExmTemplate是所有exp的模板
编写实际exp时，只需内嵌这个结构体即可, 签名如下
func (self *exp) Attack_1() (expResult exp_model.ExpResult)

Init方法用于初始化
EchoMsgWithReturn 用于实时回显，判断TmpResultCh是否不为空，否则即打印log

*/

import (
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

type ExpTemplate struct {
	Params exp_model.ExpSendParams
}

// 获取设置的http头部
func (self *ExpTemplate) GetInitExpHeaders() (headers lzhttp.Header) {
	return
}
