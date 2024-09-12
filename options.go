package Switch

import "time"

type option struct {
	Identifier *string               // 用于指定使用哪一套配置, 默认为DefaultKey
	Now        *time.Time            // 用于指定比较的时间
	Pid        *int                  // 用于指定需要判定的Pid
	Vip        *int                  // 用于指定需要判定的Pid
	Pkg        *string               // 用于指定需要判定的Pkg
	Func       func(string, *Result) // 所有的逻辑都走完了,再次可以修改结果
	IsAsk      *bool                 // 用于指定是否需要询问
}

func Option() *option {
	return &option{}
}

func (this *option) SetPid(pid int) *option {
	if this == nil {
		return nil
	}
	this.Pid = &pid
	return this
}

func (this *option) SetVip(vip int) *option {
	if this == nil {
		return nil
	}
	this.Vip = &vip
	return this
}

func (this *option) SetNow(t time.Time) *option {
	if this == nil {
		return nil
	}
	this.Now = &t
	return this
}

func (this *option) SetIsAsk(b bool) *option {
	if this == nil {
		return nil
	}
	this.IsAsk = &b
	return this
}

func (this *option) SetIdentifier(delta string) *option {
	if this == nil {
		return nil
	}
	this.Identifier = &delta
	return this
}

func (this *option) SetPkg(delta string) *option {
	if this == nil {
		return nil
	}
	this.Pkg = &delta
	return this
}

func (this *option) SetFunc(f func(string, *Result)) *option {
	if this == nil {
		return nil
	}
	this.Func = f
	return this
}

func (this *option) merge(delta *option) *option {
	if this == nil || delta == nil {
		return this
	}

	if delta.Pid != nil {
		this.Pid = delta.Pid
	}

	if delta.Vip != nil {
		this.Vip = delta.Vip
	}

	if delta.Now != nil {
		this.Now = delta.Now
	}
	if delta.IsAsk != nil {
		this.IsAsk = delta.IsAsk
	}
	if delta.Identifier != nil {
		this.Identifier = delta.Identifier
	}
	if delta.Pkg != nil {
		this.Pkg = delta.Pkg
	}

	if delta.Func != nil {
		this.Func = delta.Func
	}
	return this
}

func (this *option) Merge(deltas ...*option) *option {
	if this == nil || deltas == nil {
		return this
	}
	for _, delta := range deltas {
		this.merge(delta)
	}
	return this
}
