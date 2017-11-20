package toy2

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

// package bufio 实现了 缓存IO
// 包装了 io.Reader / io.Writer 对象，实现了 有缓存的 io.Reader / io.Writer

// bufio.Reader
// type Reader struct {
// 	buf          []byte    // 缓存
// 	rd           io.Reader // 底层 Reader
// 	r, w         int       // r: 从buf中读走的字节 w:buf中填充内容的偏移
// 	err          error
// 	lastByte     int // 最后一次读到的字节
// 	lastRuneSize int // 最后一次读取到的 rune 的大小
// }

// func NewReader(rd io.Reader) *Reader {
// 	// defaultBufSize = 4096
// 	return NewReaderSize(rd, defaultBufSize)
// }

// func newReaderSize(rd io.Reader, size int) *bufio.Reader {
// 	// 已经是 bufio.Reader 类型，且缓存大小不小于size的，直接返回
// 	b, ok := rd.(*bufio.Reader)
// 	if ok && len(b.buf) >= size {
// 		return b
// 	}
// 	if size < 4096 {
// 		size = 4096
// 	}
// 	return &bufio.Reader{
// 		buf:          make([]byte, size),
// 		rd:           rd,
// 		lastByte:     -1,
// 		lastRuneSize: -1,
// 	}
// }

// ReadSlice ReadBytes ReadString ReadLine
// ReadSlice 从输入中读取，直到遇到第一个界定符（delim）为止，返回一个指向缓冲中字节的slice
// 在下次调用读操作（read）时，这些字节会无效
// func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadSlice return a []byte -> buf
func DemoReadSlice() {
	ioReader := strings.NewReader("http://studygolang.com. \nIt is the home of gophers.\n Here is more with no new line.")
	reader := bufio.NewReader(ioReader)

	line, _ := reader.ReadSlice('\n')
	fmt.Printf("line1: %s", line)

	reader.ReadSlice('\n')

	n, _ := reader.ReadSlice('\n')

	fmt.Printf("line1: %s", line)
	fmt.Printf("line2: %s", n)
}

func DemoReadBytes() {
	ioReader := strings.NewReader("http://studygolang.com. \nIt is the home of gophers\n. Here is more with no new line.")
	reader := bufio.NewReader(ioReader)

	line, _ := reader.ReadBytes('\n')
	fmt.Printf("line1: %s", line)

	n, _ := reader.ReadBytes('\n')
	fmt.Printf("line2: %s", n)

	fmt.Printf("line1: %s", line)
}

func DemoReadString() {
	ioReader := strings.NewReader("http://studygolang.com. \nIt is the home of gophers\n. Here is more with no new line.")
	reader := bufio.NewReader(ioReader)

	line, _ := reader.ReadString('\n')
	fmt.Printf("line1: %s", line)

	n, _ := reader.ReadString('\n')
	fmt.Printf("line2: %s", n)

	fmt.Printf("line1: %s", line)
}

// ReadLine
// 逐行读取命令 -> ReadBytes('\n') , ReadString('\n')
// ReadLine 不包含行尾的\n  或者 \r\n
// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

// Peek
// 看 Reader 中有没有读取的 n 个字节
// 返回的 []byte 是 b.buf 的引用， 下次io操作后无效
// 所以 诸如 ReadSlice 这样的方法 对于多 goroutine 是不安全的
// func (b *Reader) Peek(n int) ([]byte, error)
func demoPeekLine(r *bufio.Reader) {
	line, _ := r.Peek(14)
	fmt.Printf("%s\n", line)
	time.Sleep(1)
	fmt.Printf("%s\n", line)
}

func DemoPeek() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com.\t It is the home of gophers."))
	go demoPeekLine(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}
