package Switch

func SetDefaultKey(key string) {
	default_key = key
}

func SetDefaultKeyStartTime(key string) {
	default_key_sarttime = key
}

func SetDefaultKeyEndTime(key string) {
	default_key_endtime = key
}

// SetAskFunc 设置询问函数
func SetAskFunc(f func(string, *Result, *option)) {
	ask_func = f
}

func Open(opts ...*option) (m map[string]*Result) {
	return DefaultSwitch.Open(opts...)
}

func IsOpen(acname string, opts ...*option) (r *Result) {
	return DefaultSwitch.IsOpen(acname, opts...)
}

var DefaultSwitch = &Switch[*IdentifierSwitchItem[*SwithItem]]{}
