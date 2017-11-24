package toy6

import (
	"testing"
)

// func TestXxx(t *testing.T)
// Xxx 可以使任何字母数字字符串，但是第一个字母不能是小写字母
// test 基本测试用例
// benchmark 压力测试
// example 测试控制台输出
// testMain 测试main函数

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Test_Fib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := fib(in)
	if actual != expected {
		t.Errorf("fib(%d) = %d; expected %d\n", in, actual, expected)
	}
}

// 覆盖率？table-driven
func Test_FibTableDriven(t *testing.T) {
	var fibTests = []struct {
		in       int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTests {
		actual := fib(tt.in)
		if actual != tt.expected {
			t.Errorf("fib(%d) = %d; expected %d\n", tt.in, actual, tt.expected)
		}
	}
}

// func ExampleGetFib() {
// 	r := fib(8)
// 	fmt.Println(r)
// 	// Output:
// 	// 21
// }
