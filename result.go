package golang_sdk

type Result struct {
	State int
	Msg   string
	Data  interface{}
}

func ErrorResult(state int, msg string) *Result {
	return &Result{
		State: state,
		Msg:   msg,
	}
}
