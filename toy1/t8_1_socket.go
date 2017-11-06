package toy1

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

// open -> write / read -> close
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/images/8.1.socket.png?raw=true

// TCPConn - implementation of the Conn interface for TCP network connections.
// Write
// Read

// ResolveTCPAddr - tcp4 tcp6 tcp
func demo1() (*net.TCPAddr, error) {
	addrstr := "127.0.0.1:80"
	tcpaddr, err := net.ResolveTCPAddr("tcp4", addrstr)
	if err != nil {
		return nil, err
	}
	return tcpaddr, nil
}

// CheckError format error and write to os.Stderr
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func demo2() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777")
	CheckError(err)

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)

	// io - ioutil - ReadAll
	// Reads from io.Reader until an error or EOF and returns the data it read.
	// A successful call returns err == nil not err == EOF, cause it is defined to read
	// from src until EOF not treat and EOF from Read as an error to be reported.
	result, err := ioutil.ReadAll(conn)
	CheckError(err)
	fmt.Println(string(result))

	os.Exit(0)
}

func Demo2() {
	demo2()
}

func handleClient(conn net.Conn) {
	fmt.Fprintf(os.Stdout, "hanling client", nil)
	defer conn.Close()
	conn.Write([]byte(time.Now().String()))
}

func Demo3() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777")
	CheckError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}
