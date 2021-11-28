package fs

import (
	"io"
	"net"
	"os"
)

func Ls(conn net.Conn, cmd string) {
	conn.Write([]byte(cmd + "\n"))

	defer conn.Close()

	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		return
	}
}
