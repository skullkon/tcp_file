package fs

import (
	"io"
	"log"
	"net"
	"os"
)

func GetFile(conn net.Conn, fname string, cmd string) {
	conn.Write([]byte(cmd + "\n"))
	conn.Write([]byte(fname))
	defer conn.Close()
	file, err := os.Create(fname)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Println(err)
		return
	}
}
