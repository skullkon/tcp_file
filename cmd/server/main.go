package main

import (
	"github.com/joho/godotenv"
	"github.com/skullkon/tcp_file/internal/ftp"
	"net"
)

var (
	_ = godotenv.Load()
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go ftp.HandleConn(conn)
	}

}
