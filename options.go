package Switch

import "time"

type Option struct {
	identifier       *string               // 用于指定使用哪一套配置, 默认为DefaultKey
	now              *time.Time            // 用于指定比较的时间
	pid              *int                  // 用于指定需要判定的Pid
	vip              *int                  // 用于指定需要判定的Pid
	pkg              *string               // 用于指定需要判定的Pkg
	channel          *string               // 用于指定需要判定的channel
	fix_func_finally func(string, *Result) // 所有的逻辑都走完了,对最终结果的修正
}

func Options() *Option {
	return &Option{}
}

func (this *Option) SetPid(pid int) *Option {
	if this == nil {
		return nil
	}
	this.pid = &pid
	return this
}
func (this *Option) GetPid() int {
	if this == nil || this.pid == nil {
		return 0
	}
	return *this.pid
}

func (this *Option) SetVip(vip int) *Option {
	if this == nil {
		return nil
	}
	this.vip = &vip
	return this
}

func (this *Option) GetVip() int {
	if this == nil || this.vip == nil {
		return 0
	}
	return *this.vip
}

func (this *Option) SetNow(t time.Time) *Option {
	if this == nil {
		return nil
	}
	this.now = &t
	return this
}

func (this *Option) GetNow() time.Time {
	if this == nil || this.now == nil {
		return time.Time{}
	}
	return *this.now
}

func (this *Option) SetIdentifier(delta string) *Option {
	if this == nil {
		return nil
	}
	this.identifier = &delta
	return this
}

func (this *Option) GetIdentifier() string {
	if this == nil || this.identifier == nil {
		return ""
	}
	return *this.identifier
}

func (this *Option) SetPkg(delta string) *Option {
	if this == nil {
		return nil
	}
	this.pkg = &delta
	return this
}

func (this *Option) GetPkg() string {
	if this == nil || this.pkg == nil {
		return ""
	}
	return *this.pkg
}

func (this *Option) SetChannel(delta string) *Option {
	if this == nil {
		return nil
	}
	this.channel = &delta
	return this
}

func (this *Option) GetChannel() string {
	if this == nil || this.channel == nil {
		return ""
	}
	return *this.channel
}

func (this *Option) SetFunc(f func(string, *Result)) *Option {
	if this == nil {
		return nil
	}
	this.fix_func_finally = f
	return this
}

func (this *Option) merge(delta *Option) *Option {
	if this == nil || delta == nil {
		return this
	}

	if delta.pid != nil {
		this.pid = delta.pid
	}

	if delta.vip != nil {
		this.vip = delta.vip
	}

	if delta.now != nil {
		this.now = delta.now
	}

	if delta.identifier != nil {
		this.identifier = delta.identifier
	}

	if delta.pkg != nil {
		this.pkg = delta.pkg
	}

	if delta.channel != nil {
		this.channel = delta.channel
	}

	if delta.fix_func_finally != nil {
		this.fix_func_finally = delta.fix_func_finally
	}

	return this
}

func (this *Option) merges(deltas ...*Option) *Option {
	if this == nil || deltas == nil {
		return this
	}
	for _, delta := range deltas {
		this.merge(delta)
	}
	return this
}
