package exp_templates

// 调试信息打印，开启DEBUG方可显示
func (self *ExpTemplate) EchoDebugMsg(msg string) {}

// 信息打印
func (self *ExpTemplate) EchoInfoMsg(msg string) {}

// 不换行信息打印
func (self *ExpTemplate) EchoInfoMsgWithoutReturn(msg string) {}

// 错误信息打印
func (self *ExpTemplate) EchoErrMsg(msg string) {}

// 不换行错误信息打印
func (self *ExpTemplate) EchoErrMsgWithoutReturn(msg string) {}
