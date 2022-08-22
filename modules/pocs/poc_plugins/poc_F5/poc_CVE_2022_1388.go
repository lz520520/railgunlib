package poc_F5

import (
	"github.com/lz520520/railgunlib/pkg/register/poc_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
	"github.com/lz520520/railgunlib/pkg/templates/poc_templates"
	"strings"
)

type Poc_CVE_2022_1388 struct {
	poc_templates.PocTemplate
}

func (self *Poc_CVE_2022_1388) Poc1() (pocResult poc_model.PocPerPayloadResult) {
	// 默认配置
	pocResult.Status = false
	headers := self.GetInitPocHeaders()
	resp := self.HttpGetWithoutRedirect(self.AddUri(self.Params.Target, "/mgmt/shared/authn/login"), headers)
	if resp.Err != nil {
		return
	}
	if strings.Contains(resp.Body, "resterrorresponse") &&
		strings.Contains(resp.Body, "Authorization failed") {
		pocResult.Status = true
	}
	//"{\"success\":false,\"msg\":\"Verification failure\"}"

	return
}

func init() {
	registerMsg := poc_register.PocRegisterMsg{Msg: poc_model.PocMsg{
		Name:     "F5 iControl REST API Auth Bypass",
		Author:   "lz520520",
		Time:     "2022-05-05",
		Range:    "",
		ID:       "CVE-2022-1388",
		Describe: "F5 iControl REST API Auth Bypass",
		VulType:  common.VulUnauth,
	},
		Proto: poc_model.PocHTTP,
	}
	poc_register.PocStructRegister(&Poc_CVE_2022_1388{}, registerMsg)
}
