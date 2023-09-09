package errs

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
)

type ErrorCode struct {
	code    string
	message string
	caller  string
}

func New(message string) *ErrorCode {
	_, fn, line, _ := runtime.Caller(1)
	caller := strings.Join([]string{fn, strconv.Itoa(line)}, ":")
	return &ErrorCode{code: defaultErr, message: message, caller: caller}
}

func (e *ErrorCode) Error() string {
	return e.code + ": " + e.message
}

func (e *ErrorCode) GetCode() string {
	return e.code
}

func (e *ErrorCode) GetMessage() string {
	return e.message
}

func (e *ErrorCode) Caller() string {
	return e.caller
}

func (e *ErrorCode) IsErr(target error) bool {
	//t, ok := ParseErr(target)
	//if !ok {
	//	return false
	//}
	//return e.GetCode() == t.GetCode()
	return errors.Is(target, e)
}

func ParseErr(err error) (*ErrorCode, bool) {
	var e *ErrorCode
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}
