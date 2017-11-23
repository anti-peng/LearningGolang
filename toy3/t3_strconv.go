package toy3

import (
	"fmt"
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
const theIntSize = 32 << uint(^uint(0)>>63)
const TheIntSize = theIntSize

// 整型 --> string
// FormatUint FormatInt Itoa

// uint -> string 实例 faster than i.interface
func DemoDec2Str() {

	const digits = "0123456789abcdefghijklmnopqrstuvwxyz"
	u := uint64(127) // 127

	var a [65]byte
	i := len(a) // 65

	b := uint64(10) // 10 base

	// 每次取个位
	for u >= b {
		i--                         // 65, 64, ..., 10
		a[i] = digits[uintptr(u%b)] // 取 个位
		u /= b                      // prepare for the next gen
	}
	i--
	a[i] = digits[uintptr(u)]
	fmt.Println(string(a[1:]))
}
