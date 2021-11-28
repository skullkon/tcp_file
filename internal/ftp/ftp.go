package ftp

import (
	"io"
	"net"
	"os"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	for {

		file, _ := os.Create("./file.txt")
		defer file.Close()
		_, _ = io.Copy(file, conn)
		break
	}
}
