package exp_templates

// 成功信息打印
func (self *ExpTemplate) EchoSuccessMsg(format string, a ...any) {}

// 告警信息打印
func (self *ExpTemplate) EchoWarnMsg(format string, a ...any) {}

// 普通信息打印
func (self *ExpTemplate) EchoInfoMsg(format string, a ...any) {}

// 错误信息打印
func (self *ExpTemplate) EchoErrMsg(format string, a ...any) {}

// 调试信息打印，开启DEBUG方可显示
func (self *ExpTemplate) EchoDebugMsg(format string, a ...any) {}

// 不换行信息打印
func (self *ExpTemplate) EchoInfoMsgWithoutReturn(format string, a ...any) {}

// 不换行错误信息打印
func (self *ExpTemplate) EchoErrMsgWithoutReturn(format string, a ...any) {}
