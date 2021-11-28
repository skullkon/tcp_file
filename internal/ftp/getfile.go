package ftp

import (
	"github.com/google/uuid"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	root = "./fileServer/"
)

func GetFile(conn net.Conn, fName string) {

	file, err := os.Create(root + uuid.New().String() + "." + strings.Split(fName, ".")[2])
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
	conn.Close()
}
