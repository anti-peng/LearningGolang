package toy3

import (
	"fmt"
	"unicode/utf8"
)

// package unicode

// unicode 1 2
// utf8 1 3
// utf16 2 4

// in golang, "xxxx" is utf8 string, 'a' 'å' is unicode : rune

// package utf8
// Valid ValidRune ValidString

func DemoUTF8() {
	ch := '中'
	fmt.Println(utf8.ValidRune(ch))
	fmt.Println(utf8.RuneLen(ch))

	// var bs []byte
	// bs := []byte("hello gophers!")
	bs := []byte("1234")
	effected := utf8.EncodeRune(bs, ch)
	fmt.Printf("Effected: %d, we get %s\n", effected, bs)
}
