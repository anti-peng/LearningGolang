package toy2

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// io
// io/ioutil
// fmt
// bufio

// Read reads up to len(p) into p. It returns the number of bytes read(0 <= n <= p)
// and any error. Even if n < len(p), 它也会在调用过程中使用p的全部作为暂存空间。
// If some data is available but not len(p) bytes, Read 会照例返回可用的数据，
// 而不是等待更多数据
// 当成功读取了 n > 0 个字节后遇到一个错误或者EOF，READ会返回读取的字节数。
// 它会从相同的调用中返回非nil的错误或者从随后的调用中返回错误，同时n = 0；
// 也就是，READER在输入流结束会返回读取的非零字节数，同时返回的ERR不是EOF就是nil，并且
// 下一个READ返回的一定是 0, EOF
// 换一个说法，READ 方法返回错误的时候不代表没有读取到数据，调用者应该处理返回的任何数据，
// 之后才处理可能的错误
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// toy2ReadFrom 可以从任意地方读取数据，只要来源实现了 io.Reader 接口
// data, err := toy2ReadFrom(os.Stdin, 11)
// data, err := toy2ReadFrom(file, 9)
// data, err := toy2ReadFrom(strings.NewReader("from string"), 12)
func toy2ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func genReaderMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}

// Readerexample show how to implements io.Reader interface in different ways
func Readerexample() {
FOREND:
	for {
		genReaderMenu()

		// get terminal inputs
		var ch string
		fmt.Scanln(&ch)

		var (
			data []byte
			err  error
		)

		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，以回车结束：")
			data, err = toy2ReadFrom(os.Stdin, 11)
		case "2":
			file, err := os.Open("/Users/fry/GreatLD/gogogo/src/LearningGolang/README.md")
			defer file.Close()
			if err != nil {
				fmt.Println("打开文件失败:", err)
				continue
			}
			data, err = toy2ReadFrom(file, 9)
		case "3":
			data, err = toy2ReadFrom(strings.NewReader("from string haha"), 12)
		case "4":
			fmt.Println("wait")
		case "b":
			fmt.Println("返回上级菜单")
			break FOREND
		case "q":
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败")
		} else {
			fmt.Printf("读取到的数据是: %s\n", data)
		}
	}
}
