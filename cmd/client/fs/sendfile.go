package fs

import (
	"io"
	"net"
	"os"
)

func SendFile(conn net.Conn, fname string, cmd string) {
	conn.Write([]byte(cmd + "\n"))
	conn.Write([]byte(fname))
	file, err := os.Open(fname)
	if err != nil {

		return
	}
	defer conn.Close()
	_, err = io.Copy(conn, file)
	if err != nil {
		return
	}
}
