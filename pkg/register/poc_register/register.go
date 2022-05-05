package poc_register

import (
	"github.com/lz520520/railgunlib/pkg/templates/poc_model"
)

type PocRegisterMsg struct {
	Msg   poc_model.PocMsg
	Proto poc_model.PocProto
}

func PocStructRegister(s interface{}, registerMsg PocRegisterMsg) {
}
