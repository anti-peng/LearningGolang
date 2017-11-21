package toy3

import (
	"bytes"
	"fmt"
	"strings"
)

type theContains func(s, substr string) bool     // substr 在 s 中，返回 true
type theContainsAny func(s, chars string) bool   // chars 中任何一个 unicode代码点在s中，返回true
type theContainsRune func(s string, r rune) bool // unicode代码点 r 在 s 中，返回 true

func demoContainsGroup() {
	// 第二个参数 chars 中任意一个字符（Unicode Code Point）如果在第一个参数 s 中存在，则返回true。
	fmt.Println(strings.ContainsAny("team", "i"))         // false
	fmt.Println(strings.ContainsAny("failure", "u & i"))  // true
	fmt.Println(strings.ContainsAny("in failure", "s g")) // true
	fmt.Println(strings.ContainsAny("foo", ""))           // false
	fmt.Println(strings.ContainsAny("", ""))              // false
}

// substr 出现次数
// 朴素匹配算法
// KMP算法
// Rabin-Karp算法
type theCount func(s, sep string) int

// Boyer-Moore算法

// string -> []substr
// Fields FieldsFunc
// Fields 用一个或多个连续的空格 分割 s, 如果 s 只包含空格，返回 []string len=0 空格的定义是 unicode.IsSpace
type theFields func(s string) []string

// FieldsFunc 满足 f(c) 返回 true
// strings.FieldsFunc("bar foo baz ", unicode.IsSpace)
func myFieldsFunc(r rune) bool {
	return byte(r) == '*'
}
func DemoFieldsFunc() {
	fmt.Println(strings.FieldsFunc("hello*gophers!", myFieldsFunc))
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
}

// Split SplitAfter SplitN SplitAfterN
// func Split(s, sep string) []string {return genSplit(s, sep, 0, -1)}
// 通过 sep 分割，返回 []string; Split("abc", "") -> [a b c]

// HasPrefix
func theHasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
func theHasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):len(s)] == suffix
}

// str or char 出现在 字符串中的位置
// 在 s 中查找 sep 的第一次出现 返回位置
// func Index(s, sep string) int
// chars中任意一个unicode代码点在s中首次出现的位置
// func IndexAny(s, chars string) int
// Unicode 代码点 r 在 s 中第一次出现的位置
// func IndexRune(s string, r rune) int

func DemoIndexFunc() {
	fmt.Printf("%d\n", strings.IndexFunc("hello gophers!", func(r rune) bool {
		return r > 'h'
	}))
}

// Join
// func Join(a []string, sep string) string
func myJoin(str []string, sep string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buf := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	return buf.String()
}

// source code, avoid import package bytes
func srcJoin(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	// 计算返回的字符串的总长度
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func DemoSrcJoin() {
	a := []string{"hello", "gophers", "!"}
	str := srcJoin(a, "*")
	fmt.Println(str)
}

// Repeat
// func Repeat(s string, count int) string

// Replace
// func Replace(s, old, new string, n int) string
// n < 0 全部替换；

// Replacer
// func NewReplacer(oldnew ...string) *Replacer
func demoReplacer() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
}

// Replacer 还有一个方法
// 将替换结果写入到 io.Writer 中
// func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

// Reader 类型
// impl io.Reader
// 实现了 io.Reader io.ReaderAt io.Seeker io.WriterTo io.ByteReader io.ByteScanner io.RuneReader
// io.RuneScanner
type theReader struct {
	s        string // 数据源
	i        int    // 当前索引位置
	prevRune int    // index of previous rune; or < 0 （前一个读取的rune索引位置）
}

// strings.Reader 只有一个实例化方法
// func NewReader(s string) *Reader
