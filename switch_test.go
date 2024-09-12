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
	SetAskFunc(func(acname string, r *Result, opt *option) {
		fmt.Printf("ask: %s,%+v,%+v\n", acname, r, opt)
		// r.Is = false
	})
	DefaultSwitch.Load(byte)
	cde := m.Run()
	os.Exit(cde)
}

func Test_test(t *testing.T) {
	v := DefaultSwitch.IsOpen("cdkey", Option().SetPkg("SGC").SetPid(998).SetIsAsk(true).SetFunc(func(acname string, r *Result) {
		// r.Is = true
	}))
	t.Logf("%+v", v)
}
