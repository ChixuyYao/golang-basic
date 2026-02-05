package ginx

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
	Meta any    `json:"meta"`
}
