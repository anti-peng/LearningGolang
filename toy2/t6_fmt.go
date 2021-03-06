package toy2

import "fmt"
import "strconv"

// %v 默认格式
// %+v 结构体带字段名
// %#v 相应值得Go语法表示
// %T 值类型的Go语法表示
// %% %

// %t 布尔占位符 true / false

// 整数占位符
// %b 二进制表示
// %c Unicode码点所表示的字符
// %d 十进制
// %o 八进制
// %q quote 单引号围绕的字符字面量 由Go语法安全的转义
// %x 十六进制 a-f
// %X 十六进制 A-F
// %U Unicode U+1234 等同于 U+%04X

// 浮点和复数的组成部分（实部和虚部）
// %b 无小数部分的，指数为2的幂的科学计数法 -12345467p-78
// %e 科学计数法 -1234.456e+78
// %E -1234.456E+78
// %f 有小数点的无指数 123.456
// %g		根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%g", 10.20)							10.2
// %G		根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%G", 10.20+2i)						(10.2+2i)

// %s		输出字符串表示（string类型或[]byte)							Printf("%s", []byte("Go语言中文网"))		Go语言中文网
// %q		双引号围绕的字符串，由Go语法安全地转义							Printf("%q", "Go语言中文网")				"Go语言中文网"
// %x		十六进制，小写字母，每字节两个字符								Printf("%x", "golang")						676f6c616e67
// %X		十六进制，大写字母，每字节两个字符								Printf("%X", "golang")						676F6C616E67

// %p 十六进制表示的 指针占位符 前缀 0x

// 其他标记
// + 总打印数值的正负号 对于 %+q 保证只输出ASCII编码的字符 %+q '中' => "\u4e2d\u6587"
// - 在右侧而非左侧填充空格
// # 备用格式 八进制前导0 (%#o) 为十六进制添加前导0x (%#x) => Printf("%#U", '中') U+4E2D '中'
// 0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始
// （即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的
// Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
// ' '		（空格）为数值中省略的正负号留出空白（% d）；
// 以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
// 0		填充前导的0而非空格；对于数字，这会将填充移到正负号之后

// Scanning

// Fprint Fprintf Fprintln Sprint Sprintf Sprintln Print Printf Println
// Fpxxx writer -> writer
// Prxxx interface{} -> os.Stdout
// Sprxx interface{} -> string

// type Stringer interface {
// 	String() string
// }

// type Formatter interface {
// 	Format(f State, c rune)
// }
func demoFormat(f fmt.State, c rune) {
	type person struct {
		name string
		age  int
	}

	var p = &person{"fry", 11}

	if c == 'L' {
		f.Write([]byte(p.name + " " + strconv.Itoa(p.age)))
		f.Write([]byte("person has two fields"))
	} else {
		f.Write([]byte(fmt.Sprintln(p.name + "-" + strconv.Itoa(p.age))))
	}
}
