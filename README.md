#### 暴露的API方法
```golang
// 同一个模块,可能有多套配置,指定一个默认配置,比如配置文件的default
func SetDefaultKey(key string)

// 设置时间注入的key值,如果希望把开始时间注入到返回结果中,需要特殊指定
func SetDefaultKeyStartTime(string)

// 设置时间注入的key值,如果希望把结束时间注入到返回结果中,需要特殊指定,一般用于结束倒计时之类的
func SetDefaultKeyEndTime(string)

// SetAskFunc 设置询问函数,包括远程rpc调用就在这里
func SetAskFunc(func(string, ISwitchItem, *Result, *Option_))

// 设置一个自行实现的,默认实现不理想的话,可以替换
func SetDefaultSwitch(s ISwitch)

// 配置的所有开关
func Open(...*Option_) map[string]*Result

// 单个模块的开关情况,第一个参数是活动名称. eg: pingfen
func IsOpen(string, ...*Option_) *Result

// 加载配置文件
func Load([]byte) error
```

#### 配置文件
```yaml
pingfen: # 活动名称
    default: #指定用哪一套配置
        Open: true #是否开启
        StartTime: 2023-07-14T17:00:00.000Z #开始时间, 为空表示不限制
        EndTime: 2023-07-31T17:00:00.000Z #结束时间, 为空表示不限制
        Pids: # 限制访问的进程pid, 为空表示不限制
        PidTails: # 限制访问的pid尾号, 为空表示不限制
        Pkgs: # 限制访问的软件包, 为空表示不限制
        Vips: #限制访问的vip, 为空表示不限制
        Meta: null #透传一些元数据
        Server: "" #如果有远程询问,指定询问的服务
        Module: "" #如果有远程询问,指定询问的模块
        Method: "" #如果有远程询问,指定询问的方法
```
