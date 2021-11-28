package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	for {
		file, err := os.Open("./hello.txt")
		if err != nil {
			return
		}
		_, err = io.Copy(conn, file)
		break
	}
}
