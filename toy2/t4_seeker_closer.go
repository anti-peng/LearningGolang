package toy2

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Seek 设置下一次 Read / Write 的偏移量为 offset
// whence： 0 1 2 -> 0 I N 偏移策略
// Seek 返回新的偏移量和错误，if exists
// type Seeker interface {
// 	Seek(offset int64, whence int) (ret int64, err error)
// }

func demoSeek() {
	reader := strings.NewReader("0123456789")
	// offset, err := reader.Seek(-2, os.SEEK_END)
	// offset, err := reader.Seek(8, os.SEEK_SET)
	content := make([]byte, 8)
	allRead, err := reader.Read(content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("allRead: %d\n", allRead)

	offset, err := reader.Seek(0, os.SEEK_CUR)
	if err != nil {
		panic(err)
	}
	// r, _, _ := reader.ReadRune()
	// fmt.Println("offset now: " + strconv.FormatInt(offset, 10))
	// fmt.Printf("%c\n", r)

	tailContent := make([]byte, 10-offset)
	_, err = reader.Read(tailContent)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tailReadContent: %s\n", string(tailContent))
}

// 关闭数据流
// defer readerOrWriter.Close() should be placed after error checker
// cause readerOrWriter might be nil, and nil doesn't implements
// interface Closer
// type Closer interface {
// 	Close() error
// }

// Others
// Read or write one byte
// 标准库中有些类型实现了 io.ByteWriter / io.ByteReader
// bufio.Reader/Writer
// bytes.Buffer
// bytes.Reader
// strings.Reader
// type ByteReader interface {
// 	ByteRead() (c byte, err error)
// }
// type ByteWriter interface {
// 	ByteWrite(b byte) error
// }

// 一般不会使用bytes.Buffer 一次读取或者写入一个字节
// 标准库 encoding/binary 中 ReadVarint 需要
// 标准库 image/jpeg Encode -> implements ByteWriter
func demoByteReadWrite() {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	if err := buffer.WriteByte(ch); err != nil {
		panic(err)
	}
	fmt.Println("Successfully write one byte! Ready to read it...")
	newCh, _ := buffer.ReadByte()
	fmt.Printf("byte read: %c\n", newCh)
}

// 继承了 ByteReader 接口
//  UnreadByte 将上一次 ReadByte 的字节还原，使得在此调用 ReadByte 返回的结果和上一次一样
// 调用 UnreadByte 之前 必须调用了 ReadByte，
// 而且不能连续调用 UnreadByte
// type ByteScanner interface {
// 	ByteReader
// 	UnreadByte() error
// }
func readOneByte(buf *bytes.Buffer) {
	b, err := buf.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("byte read %s\n", string(b))
}
func demoByteScanner() {
	buffer := bytes.NewBuffer([]byte{'a', 'b', 'c'})

	readOneByte(buffer)
	readOneByte(buffer)

	if err := buffer.UnreadByte(); err != nil {
		panic(err)
	}

	readOneByte(buffer)
	readOneByte(buffer)
}

// SectionReader
// 一个 struct 没有任何导出字段 实现了 Read Seek 和 ReadAt
// 内嵌 ReaderAt
// type SectionReader struct {
// 	r     ReaderAt
// 	base  int64
// 	off   int64
// 	limit int64
// }

// LimitedReader
// 从 R 读取但将返回数据量限制为 N 个字节
// type LimitedReader struct {
// 	R Reader
// 	N int64
// }

// package io 中，LimitReader 的实现是调用 LimitedReader
// func LimitReader(r Reader, n int64) Reader {
// 	return &LimitedReader{r, n}
// }
func demoLimitReader() {
	reader := strings.NewReader("this is an LimitReader Example")
	limitReader := &io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		limitReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}
	fmt.Print("\n")
}

// PipeReader 和 PipeWriter
// implements interface io.Reader & io.Closer

// Read - get data from channel, will block until 写入端开始写入数据
// 或者写入端关闭了。如果写入端关闭时带上error（call CloseWithError) 该
// 方法返回的 err 就是写入端传递的error，否则 err 为 EOF

// Write - 写数据到管道中 会堵塞直到管道读取端读完所有数据或者读取端关闭了。
// 读取端错误 ErrClosedPipe / 读取端传递的error

func demoPipeWrite(pw *io.PipeWriter) {
	var (
		i   = 0
		err error
		n   int
	)

	data := []byte("你好 golang")

	for n, err = pw.Write(data); err == nil; _, err = pw.Write(data) {
		i++
		if i == 3 {
			pw.CloseWithError(errors.New("输出3次后结束"))
		}
	}
	fmt.Println("close 后输出的字节数： ", n, " error: ", err)
}

func demoPipeReader(pr *io.PipeReader) {
	var (
		err error
		n   int
	)

	data := make([]byte, 1024)

	for n, err = pr.Read(data); err == nil; _, err = pr.Read(data) {
		fmt.Printf("%s\n", data[:n])
	}
	fmt.Println("writer closeWithError: ", err)
}

func demoPipe() {
	pr, pw := io.Pipe()
	go demoPipeWrite(pw)
	go demoPipeReader(pr)
	time.Sleep(time.Second * 5)
}

// Copy & CopyN
// func Copy(dst Writer, src Reader) (written int64, err error)
// copy src -> dst, until src get EOF or error;
// copy 从 src 读取直到 EOF为止
// copy 实现了 ReadFrom
// dst.ReadFrom(src)
// func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

// ReadAtLeast
// func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
// r -> buf, 直到读了最少 min 个字节为止。返回复制的字节数
// 如果读取的字节较少，会返回一个 error
// 1) 没有读取到字节，EOF
// 2）一个 EOF 发生在读取了少于 min 个字节之后，ErrUnexpectedEOF
// 3）min > len(buf) ErrShortBuffer
// 对于返回值，仅当 err == nil 才会 n >= min

// ReadFull
// func ReadFull(r Reader, buf []byte) (n int, err error)
// r -> buf
// 1) 没有读取到字节 EOF
// 2) not full-filled read ErrUnexpectedEOF
// 3) only if err == nil, n == len(buf)

// WriteString
// func WriteString(w Writer, s string) (n int, err error)

// MultiReader / MultiWriter
// func MultiReader(readers ...Reader) Reader
// func MultiWriter(writers ...Writer) Writer
// 接收多个 Reader 或 Writer，返回一个 Reader 或 Writer
// 操作多个 Reader、Writer 就像操作一个
// type multiReader struct {
// 	readers []Reader
// }
// type multiWriter struct {
// 	writers []Writer
// }
func demoMultiReader() {
	readers := []io.Reader{
		bytes.NewBufferString("from bytes reader"),
		strings.NewReader("from strings reader"),
	}

	mr := io.MultiReader(readers...)
	data := make([]byte, 1024)
	var (
		err error
		n   int
	)

	for err != io.EOF {
		tmp := make([]byte, 32)
		n, err = mr.Read(tmp)
		if err == nil {
			data = append(data, tmp[:n]...)
			data = append(data, '\n')
		} else {
			if err != io.EOF {
				panic(err)
			}
		}
	}
	fmt.Printf("%s", data)
}

func demoMultiWriter() {
	file, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writers := []io.Writer{
		file,
		os.Stdout,
	}

	writer := io.MultiWriter(writers...)
	writer.Write([]byte("你好 golang\n"))
}

// TeeReader
// func TeeReader(r Reader, w Writer) Reader
// 返回一个 Reader， r -> w, 没有缓存

func demoTeeReader() {
	reader := io.TeeReader(strings.NewReader("你好 golang\n"), os.Stdout)
	reader.Read(make([]byte, 1024))
}
