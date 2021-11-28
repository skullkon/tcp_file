package fs

import (
	"io"
	"net"
	"os"
	"strconv"
)

func SendFile(conn net.Conn, fname string, cmd string) {
	fi, err := os.Stat("./" + fname)
	conn.Write([]byte(cmd + "\n"))
	conn.Write([]byte(fname))
	conn.Write([]byte("\n" + strconv.FormatInt(fi.Size(), 10)))

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
