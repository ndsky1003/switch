package Switch

import (
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

var (
	default_key          = "default"
	default_key_sarttime = "Start"
	default_key_endtime  = "End"
)

var ask_func func(string, *Result, *option)

type ISwitchItem interface {
	IsOpen(acname string, opts ...*option) (r *Result)
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

func (this *SwithItem) IsOpen(acname string, opts ...*option) (r *Result) {
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
	opt := Option().Merge(opts...)

	r.Is = this.Open
	var now time.Time
	if opt.Now != nil {
		now = *opt.Now
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

	var pid int
	if opt.Pid != nil {
		pid = *opt.Pid
	}

	if r.Is && pid != 0 && len(this.Pids) > 0 {
		r.Is = lo.Contains(this.Pids, pid)
	}

	if r.Is && opt.Vip != nil && len(this.Vips) > 0 {
		r.Is = lo.Contains(this.Vips, *opt.Vip)
	}

	if r.Is && pid != 0 && len(this.PidTails) > 0 {
		pidtail := pid % 10
		r.Is = lo.Contains(this.PidTails, pidtail)
	}

	if r.Is && opt.Pkg != nil && len(this.Pkgs) > 0 {
		r.Is = lo.Contains(this.Pkgs, *opt.Pkg)
	}

	if r.Is && ask_func != nil && opt.IsAsk != nil && *opt.IsAsk {
		ask_func(acname, r, opt)
	}
	if opt.Func != nil {
		opt.Func(acname, r)
	}
	return
}

type IIdentifierSwitchItem interface {
	IsOpen(string, ...*option) (r *Result)
}

type IdentifierSwitchItem[T ISwitchItem] map[string]T // 一般是包分包

func (this *IdentifierSwitchItem[T]) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%+v", *this)
}

func (this *IdentifierSwitchItem[T]) IsOpen(acname string, opts ...*option) (r *Result) {
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
	opt := Option().Merge(opts...)
	var identifier string
	if opt.Identifier != nil {
		identifier = *opt.Identifier
	} else {
		identifier = default_key
	}
	if v, ok := (*this)[identifier]; ok {
		r = v.IsOpen(acname, opts...)
	}
	return
}

var rwl sync.RWMutex //protect under
type Switch[T IIdentifierSwitchItem] map[string]T

/*
返回配置的所有开关
*/
func (this *Switch[T]) Open(opts ...*option) (m map[string]*Result) {
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
func (this *Switch[T]) IsOpen(acname string, opts ...*option) (r *Result) {
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

func (this *Switch[T]) Load(buf []byte) (err error) {
	var tmp Switch[T]
	if err = yaml.Unmarshal(buf, &tmp); err != nil {
		return
	}
	rwl.Lock()
	defer rwl.Unlock()
	*this = tmp
	return
}
