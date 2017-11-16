package toy2

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
