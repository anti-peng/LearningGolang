package toy2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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

	// n, _ := reader.ReadSlice('\n')

	fmt.Printf("line1: %s", line)
	// fmt.Printf("line2: %s", n)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DemoGitHub() {
	var err error
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. It is the home of gophers."))
	line, err := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	checkErr(err)

	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, err := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
	checkErr(err)
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
// returns the next n bytes without advancing the reader
// 看 Reader 中有没有读取的 n 个字节
// 返回的 []byte 是 b.buf 的引用， 下次io操作后无效
// 所以 诸如 ReadSlice 这样的方法 对于多 goroutine 是不安全的
// func (b *Reader) Peek(n int) ([]byte, error)

func DemoPeek() {
	reader := bufio.NewReader(strings.NewReader("hello gophers!"))

	bs, err := reader.ReadBytes(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bs)

	bs, err = reader.Peek(8)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bs)
}

// ... others

// Scanner 类型和方法
// Scanner provides a convenient interface for reading data such as a file of newline-delimited
// lines of text.
// 更容易的处理 按行读入 空格分隔 等等

type theScanner struct {
	r            io.Reader
	split        bufio.SplitFunc // function to split the tokens
	maxTokenSize int             // maximum size of a token
	token        []byte          // last token returned by split
	buf          []byte          // buffer used as argument to split
	start        int             // first non-processed byte in buf
	end          int             // end of data in buf
	err          error           // sticky error
}

// SplitFunc 输入行为控制
// 进行分词，并返回未处理的数据中的第一个 token
type theSplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func DemoScanWords() {
	scanner := bufio.NewScanner(strings.NewReader("hello gophers! 中文（中文括号） E-O-F"))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Scanner 没有导出字段 需要通过 bufio 提供的实例化函数
// func theNewScanner(r io.Reader) *bufio.Scanner {
// 	return &bufio.Scanner{
// 		r:            r,
// 		split:        bufio.ScanLines,
// 		maxTokenSize: bufio.MaxScanTokenSize,
// 		buf:          make([]byte, 4096),
// 	}
// }

// bufio.Writer
type theWriter struct {
	err error
	buf []byte // 缓存
	n   int    // 当前缓存中的字节数
	wr  io.Writer
}

// init
func theNewWriter(wr io.Writer) *bufio.Writer {
	return bufio.NewWriterSize(wr, 4096)
}

// Available / Buffered
// 获取缓存中还未使用的字节数 len(buf) - n
// 获取写入当前缓存的字节数 n

// Flush
// buf -> io.Writer
// when all Write has done, call Flush -> buf -> io.Writer

// ReadFrom -> implements io.ReaderFrom
// Writer -> impl io.Writer
// WriteByte -> impl io.ByteWriter
// WriteRune -> impl by WriteByte or WriteString
// WriteString
// 这些写方法会在缓存满的会后调用 Flush

// ReadWrite
type theReadWriter struct {
	*bufio.Reader
	*bufio.Writer
}

func theNewReadWriter(r *bufio.Reader, w *bufio.Writer) *bufio.ReadWriter {
	return &bufio.ReadWriter{r, w}
}
