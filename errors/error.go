package errors

type ErrorCode struct {
	Errno  int
	ErrMsg string
}

func New() {

}
func (t *ErrorCode) Error() string {
	return t.ErrMsg
}

var (
	Success     = ErrorCode{0, "success"}
	ParamsError = ErrorCode{2, "参数错误"}
	SaveError   = ErrorCode{3, "更新失败"}
)
