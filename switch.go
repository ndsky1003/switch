package Switch

import (
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

var (
	default_key_identifier = "default"
	default_key_sarttime   = "Start"
	default_key_endtime    = "End"
)

var ask_func func(string, ISwitchItem, *Result, *Option_)

type ISwitchItem interface {
	IsOpen(acname string, opts ...*Option_) (r *Result)
}

type SwithItem struct {
	Name      string         `yaml:"Name"`
	Open      bool           `yaml:"Open"`
	StartTime time.Time      `yaml:"StartTime"`
	EndTime   time.Time      `yaml:"EndTime"`
	Pids      []int          `yaml:"Pids"`
	Vips      []int          `yaml:"Vips"`
	PidTails  []int          `yaml:"PidTails"` // 玩家ID尾号设计
	Pkgs      []string       `yaml:"Pkgs"`     // 玩家包过滤设计
	Meta      map[string]any `yaml:"Meta"`     // 配置里的字段，可以写死一些数据，但是默认指定字段Value作为调用返回值
	Server    string         `yaml:"Server"`
	Module    string         `yaml:"Module"`
	Method    string         `yaml:"Method"`
}

func (this *SwithItem) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%+v", *this)
}

func (this *SwithItem) IsOpen(acname string, opts ...*Option_) (r *Result) {
	r = &Result{
		Is: true,
	}
	if this == nil {
		return
	}
	r.Meta = this.Meta
	if r.Meta == nil {
		r.Meta = map[string]any{}
	}
	opt := Option().merges(opts...)

	r.Is = this.Open
	var now time.Time
	if opt.now != nil {
		now = *opt.now
	} else {
		now = time.Now()
	}

	if r.Is {
		if r.Is && !this.StartTime.IsZero() {
			r.Is = this.StartTime.Before(now)
			r.Meta[default_key_sarttime] = this.StartTime
		}
		if r.Is && !this.EndTime.IsZero() {
			r.Is = this.EndTime.After(now)
			r.Meta[default_key_endtime] = this.EndTime
		}
	}

	var pid, vip int
	var pkg string
	if opt.pid != nil {
		pid = *opt.pid
	}

	if opt.vip != nil {
		vip = *opt.vip
	}
	if opt.pkg != nil {
		pkg = *opt.pkg
	}

	if r.Is && pid != 0 && len(this.Pids) > 0 {
		r.Is = lo.Contains(this.Pids, pid)
	}

	if r.Is && pid != 0 && len(this.PidTails) > 0 {
		pidtail := pid % 10
		r.Is = lo.Contains(this.PidTails, pidtail)
	}

	if r.Is && len(this.Vips) > 0 {
		r.Is = lo.Contains(this.Vips, vip)
	}

	if r.Is && len(this.Pkgs) > 0 {
		r.Is = lo.Contains(this.Pkgs, pkg)
	}

	if r.Is && ask_func != nil {
		ask_func(acname, this, r, opt)
	}
	if opt.func_ != nil {
		opt.func_(acname, r)
	}
	return
}

type IIdentifierSwitchItem interface {
	IsOpen(string, ...*Option_) (r *Result)
}

type IdentifierSwitchItem[T ISwitchItem] map[string]T // 一般是包分包

func (this *IdentifierSwitchItem[T]) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%+v", *this)
}

func (this *IdentifierSwitchItem[T]) IsOpen(acname string, opts ...*Option_) (r *Result) {
	defer func() {
		if r == nil {
			r = &Result{
				Meta: map[string]any{},
			}
		}
	}()
	if this == nil {
		return
	}
	opt := Option().merges(opts...)
	var identifier string
	if opt.identifier != nil {
		identifier = *opt.identifier
	} else {
		identifier = default_key_identifier
	}
	if v, ok := (*this)[identifier]; ok {
		r = v.IsOpen(acname, opts...)
	}
	return
}

var rwl sync.RWMutex                               //protect under
type switch_[T IIdentifierSwitchItem] map[string]T //加个下划线是因为switch是关键字

/*
返回配置的所有开关
*/
func (this *switch_[T]) Open(opts ...*Option_) (m map[string]*Result) {
	m = map[string]*Result{}
	if this == nil {
		return
	}
	for k := range *this {
		m[k] = &Result{
			Meta: map[string]any{},
		}
	}
	length := len(*this)
	var wg sync.WaitGroup
	var l sync.Mutex
	arrFunc := make([]func(), 0, length)
	for k, v := range *this {
		k := k
		v := v
		arrFunc = append(arrFunc, func() {
			defer wg.Done()
			resp := v.IsOpen(k, opts...)
			l.Lock()
			defer l.Unlock()
			m[k] = resp
		})

	}
	for _, f := range arrFunc {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	return
}

/*
acname 活动名称
*/
func (this *switch_[T]) IsOpen(acname string, opts ...*Option_) (r *Result) {
	rwl.RLock()
	defer rwl.RUnlock()
	defer func() {
		if r == nil {
			r = &Result{
				Meta: map[string]any{},
			}
		}
	}()
	if this == nil {
		return
	}
	if v, ok := (*this)[acname]; ok {
		r = v.IsOpen(acname, opts...)
	}
	return
}

func (this *switch_[T]) Load(buf []byte) (err error) {
	var tmp switch_[T]
	if err = yaml.Unmarshal(buf, &tmp); err != nil {
		return
	}
	rwl.Lock()
	defer rwl.Unlock()
	*this = tmp
	return
}
