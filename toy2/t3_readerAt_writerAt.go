package toy2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadAt 从基本输入源的偏移量off处开始 Read
// 与 Read 不同之处在于：当 ReadAt 返回 n < len(p) 时，他就会返回一个非nil
// 的错误来解释为什么没有返回更多的字节。严格模式
// 即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。
// 若一些数据可用但不到 len(p) 字节，ReadAt 就会阻塞直到所有数据都可用或产生一个错误。 在这一点上 ReadAt 不同于 Read。
// type ReaderAt interface {
// 	ReadAt(p []byte, off int64) (n int, err error)
// }

func demoReadAt() {
	reader := strings.NewReader("ab你好世界")
	p := make([]byte, 2)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

// reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
func ioutilReadAllSourceCode(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.

	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()

	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func demoIoutilReadFileSourceCode(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Dont preallocate a huge buffer size
	// max <= 1e9 or else n = 0
	var n int64

	if fi, err := f.Stat(); err == nil {
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}

	return ioutilReadAllSourceCode(f, n)
}

// WriteAt

// type WriterAt interface {
// 	WriteAt(p []byte, off int64) (n int, err error)
// }

func demoWriteAt(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("golang中文社区--APPEND")
	n, err := file.WriteAt([]byte("明天上线"), 6+3*4+2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("write %d bytes at offset 20\n", n)
}

// ReaderFrom
// ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值n为读取的字节数。除 io.EOF 之外，
// 在读取过程中遇到的任何错误也会被返回
// Copy 方法在 ReaderFrom 可用的时候会优先使用它
// Any error except EOF encountered during the read is also returned.
// type ReaderFrom interface {
// 	ReadFrom(r Reader) (n int64, err error)
// }

func demoReadFrom(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

// WriterTo
//
// type WriterTo interface {
// 	WriteTo(w Writer) (n int64, err error)
// }
func demoWriteTo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bytes.NewReader([]byte("blahblah"))
	reader.WriteTo(os.Stdout)
}
