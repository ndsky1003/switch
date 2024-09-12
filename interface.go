package Switch

type ISwitch interface {
	Open(...*Option_) map[string]*Result

	IsOpen(acname string, opts ...*Option_) (r *Result)

	Load([]byte) error
}
