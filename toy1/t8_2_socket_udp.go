package toy1

import (
	"fmt"
	"net"
	"os"
	"time"
)

// DialUDP(net string, laddr, raddr)
// ListenUDP(net string, laddr *UDPAddr)
// (c *UDPConn) ReadFromUDP(b []byte)
// (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr)

func demo1UDPDial() {
	raddr, err := net.ResolveUDPAddr("udp4", ":7777")
	CheckError(err)

	conn, err := net.DialUDP("udp", nil, raddr)
	CheckError(err)

	// Implements the Conn Write method
	// Write
	_, err = conn.Write([]byte("hello from client"))
	CheckError(err)

	// Read
	chunk := make([]byte, 1024)
	n, err := conn.Read(chunk)
	CheckError(err)
	fmt.Printf("Content: %s", string(chunk[:n]))

	os.Exit(0)
}

func demo2UDPListen() {
	addr := ":7777"

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	CheckError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)

	for {
		go handleClient2(conn)
	}
}

func handleClient2(conn *net.UDPConn) {
	chunk := make([]byte, 1024)
	_, addr, err := conn.ReadFromUDP(chunk)
	if err != nil {
		return
	}

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
