package main

import (
	"net"

	"github.com/skullkon/tcp_file/internal/ftp"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go ftp.HandleConn(conn)
	}

}
