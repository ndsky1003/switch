package Switch

import "fmt"

type Result struct {
	Meta  map[string]any `json:"Meta"`
	Value any            `json:"Value"`
	Is    bool           `json:"Is"` //是否开启
}

func (this *Result) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%+v", *this)
}
