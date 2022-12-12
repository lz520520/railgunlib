package exp_register

import (
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
)

// 注册信息结构体
type ExpRegisterMsg struct {
	Msg        exp_model.ExpMsg                  // 漏洞信息
	SubOptions map[string]exp_model.ExpSubOption // 扩展信息
	AliasMap   map[string]string                 // 别名，map[string][string]{"1": "KeyCheck", "2": "RCE"}
}

// exp结构体注册
func ExpStructRegister(s interface{}, registerMsg ExpRegisterMsg) {
}

// 解析msg字符串生成ExpMsg
// Deprecated
func ExpMsgParser(msg string) exp_model.ExpMsg {
	return exp_model.ExpMsg{}
}
