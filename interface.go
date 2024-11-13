package Switch

type ISwitch interface {
	//所有配置文件里的开关情况
	Open(...*Option) map[string]*Result

	//指定活动的开关情况
	IsOpen(acname string, opts ...*Option) (r *Result)

	//加载活动配置
	Load([]byte) error
}
