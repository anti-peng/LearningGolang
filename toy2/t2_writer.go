package toy2

// Write
// Write 将 len(p)个字节从p中写入到基本数据流中
// Write 返回从p中被写入的字节数n(0<=n<=len(p))， 以及任何遇到的引起写入提前停止的错误
// 若 Write 返回的 n < len(p) 它就必须返回一个非 nil 的错误
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
// func Println(a ...interface{}) (n int, err error)
// Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// NewFile return a File with the give file descriptor and name
// 文件描述符 0input STDIN_FILENO 1output STDOUT_FILENO 2error STDERR_FILENO

// os.File 同时实现了 io.Reader 和 io.Writer
// strings.Reader 实现了 io.Reader
// bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// bytes.Buffer 同时实现了 io.Reader 和 io.Writer
// bytes.Reader 实现了 io.Reader
// compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
// crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
// encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// mime/multipart.Part 实现了 io.Reader
// net/conn 分别实现了 io.Reader 和 io.Writer(Conn接口定义了Read/Write)

// 从接口名称很容易猜到，一般地， Go 中接口的命名约定：接口名以 er 结尾。注意，这里并非强行要求，你完全可以不以 er 结尾。标准库中有些接口也不是以 er 结尾的。
