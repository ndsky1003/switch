package Switch

import "fmt"

// type Result1 struct {
// 	Is    bool
// 	Value any
// }

type Result struct {
	Meta map[string]any `json:"Meta"`
	Is   bool           `json:"Is"`
}

func (this *Result) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%+v", *this)
}
