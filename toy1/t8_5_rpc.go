package toy1

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
	"os"
)

// Client - client handler - net|
// Server - server handler - net|

// Go RPC - TCP HTTP JSONRPC
// Go RPC 仅支持 GO 开发的服务器与客户端之间的交互，因为在内部它们采用了 Gob 来编码

// Go RPC
// - 必须有两个导出类型的参数
// - 第一个参数是接收的参数，第二个是返回给客户端的参数；第二个参数必须是指针类型
// - 函数要有一个返回值 error
// - eg: func(t *T) MethodName(argv T1, replyArgv *T2) error
// T T1 T2 必须能被 encoding/gob 编码解码

// encoding/gob
// - func Register - records a type
// - func RegisterName - records a type using given name
// - Decoder / Encoder - struct
// - GobDecoder / GobEncoder - interface
// - CommonType - a struct holds name and typeId

// Server

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func server() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Client

func client() {
	cli, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		fmt.Println("DialHTTP: " + err.Error())
	}

	// Synchronous call
	args := Args{17, 8}
	var reply int

	if err := cli.Call("Arith.Multiply", args, &reply); err != nil {
		fmt.Println("Call Arith.Multiply: " + err.Error())
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient

	if err := cli.Call("Arith.Divide", args, &reply); err != nil {
		fmt.Println("Call Arith.Divide: " + err.Error())
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
