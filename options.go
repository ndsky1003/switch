package Switch

import "time"

type Option_ struct {
	identifier *string               // 用于指定使用哪一套配置, 默认为DefaultKey
	now        *time.Time            // 用于指定比较的时间
	pid        *int                  // 用于指定需要判定的Pid
	vip        *int                  // 用于指定需要判定的Pid
	pkg        *string               // 用于指定需要判定的Pkg
	func_      func(string, *Result) // 所有的逻辑都走完了,再次可以修改结果
}

func Option() *Option_ {
	return &Option_{}
}

func (this *Option_) SetPid(pid int) *Option_ {
	if this == nil {
		return nil
	}
	this.pid = &pid
	return this
}
func (this *Option_) GetPid() int {
	if this == nil || this.pid == nil {
		return 0
	}
	return *this.pid
}

func (this *Option_) SetVip(vip int) *Option_ {
	if this == nil {
		return nil
	}
	this.vip = &vip
	return this
}

func (this *Option_) GetVip() int {
	if this == nil || this.vip == nil {
		return 0
	}
	return *this.vip
}

func (this *Option_) SetNow(t time.Time) *Option_ {
	if this == nil {
		return nil
	}
	this.now = &t
	return this
}

func (this *Option_) GetNow() time.Time {
	if this == nil || this.now == nil {
		return time.Time{}
	}
	return *this.now
}

func (this *Option_) SetIdentifier(delta string) *Option_ {
	if this == nil {
		return nil
	}
	this.identifier = &delta
	return this
}

func (this *Option_) GetIdentifier() string {
	if this == nil || this.identifier == nil {
		return ""
	}
	return *this.identifier
}

func (this *Option_) SetPkg(delta string) *Option_ {
	if this == nil {
		return nil
	}
	this.pkg = &delta
	return this
}

func (this *Option_) GetPkg() string {
	if this == nil || this.pkg == nil {
		return ""
	}
	return *this.pkg
}

func (this *Option_) SetFunc(f func(string, *Result)) *Option_ {
	if this == nil {
		return nil
	}
	this.func_ = f
	return this
}

func (this *Option_) merge(delta *Option_) *Option_ {
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

	if delta.func_ != nil {
		this.func_ = delta.func_
	}

	return this
}

func (this *Option_) merges(deltas ...*Option_) *Option_ {
	if this == nil || deltas == nil {
		return this
	}
	for _, delta := range deltas {
		this.merge(delta)
	}
	return this
}
