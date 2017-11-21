package toy3

import (
	"strconv"
)

// two error types:
// ErrRange: 128 -> int8
// ErrSyntax: "" -> int
// 以上两种错误通过NumError类型构造返回
type theNumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the error (ErrRange / ErrSyntax)
}

// func (e *NumError) Error() string {

// }
func theSyntaxError(fn, str string) *strconv.NumError {
	return &strconv.NumError{fn, str, strconv.ErrSyntax}
}
func theRangeError(fn, str string) *strconv.NumError {
	return &strconv.NumError{fn, str, strconv.ErrRange}
}

// string <--> 整型
// ParseInt ParseUint Atoi
// ParseInt(string, 10, 0)
// int uint -> x86: 4byte  x64: 8byte
