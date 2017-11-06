package toy1

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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

// func handleClient(conn net.Conn) {
func handleClient(conn net.TCPConn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	conn.SetWriteDeadline(time.Now().Add(2 * time.Minute))
	// conn.SetKeepAlive(true)
	// conn.SetKeepAlivePeriod(d time.Duration)
	// func (t *net.TCPConn) SetKeepAlive(keepalive bool)
	// 这里是 net.Conn 不是 net.TCPConn

	defer conn.Close()
	request := make([]byte, 1024) // set maxium request length to 1024B to prevent flood attack

	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if readLen == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 1024) // clear last read content
	}
}

func demo2() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777")
	CheckError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	CheckError(err)

	for {
		// conn, err := listener.Accept()
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}
