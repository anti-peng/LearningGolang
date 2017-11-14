package toy2

import (
	"fmt"
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
	reader := strings.NewReader("你好世界")
	p := make([]byte, 2)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}
