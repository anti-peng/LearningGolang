package toy2

import (
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
