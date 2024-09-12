package Switch

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	byte, err := os.ReadFile("./switch.yaml")
	if err != nil {
		panic(err)
	}
	SetAskFunc(func(acname string, r *Result, opt *Option_) {
		fmt.Printf("ask: %s,%+v,%+v\n", acname, r, opt)
		// r.Is = false
	})
	default_switch.Load(byte)
	cde := m.Run()
	os.Exit(cde)
}

func Test_test(t *testing.T) {
	v := default_switch.IsOpen("cdkey", Option().SetPkg("SGC").SetPid(998).SetFunc(func(acname string, r *Result) {
		// r.Is = true
	}))
	t.Logf("%+v", v)
}
