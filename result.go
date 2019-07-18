package golang_sdk

type Result struct {
	State int `json:"state"`
	Msg   string `json:"msg"`
	Data  interface{} `json:"data"`
}

func ErrorResult(state int, msg string) *Result {
	return &Result{
		State: state,
		Msg:   msg,
	}
}
