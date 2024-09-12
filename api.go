package Switch

// 同一个模块,可能有多套配置,指定一个默认配置
func SetDefaultKey(key string) {
	default_key_identifier = key
}

// 设置时间注入的key值
func SetDefaultKeyStartTime(key string) {
	default_key_sarttime = key
}

// 设置时间注入的key值
func SetDefaultKeyEndTime(key string) {
	default_key_endtime = key
}

// SetAskFunc 设置询问函数,包括远程rpc调用就在这里
func SetAskFunc(f func(string, ISwitchItem, *Result, *Option_)) {
	ask_func = f
}

// 设置一个自行实现的
func SetDefaultSwitch(s ISwitch) {
	default_switch = s
}

// 配置的所有开关
func Open(opts ...*Option_) (m map[string]*Result) {
	return default_switch.Open(opts...)
}

// 单个模块的开关情况
func IsOpen(acname string, opts ...*Option_) (r *Result) {
	return default_switch.IsOpen(acname, opts...)
}

// 加载配置文件
func Load(buf []byte) (err error) {
	return default_switch.Load(buf)
}

var default_switch ISwitch = &switch_[*IdentifierSwitchItem[*SwithItem]]{}
