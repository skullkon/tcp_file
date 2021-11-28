package ftp

import (
	"io"
	"log"
	"net"
	"os"
)

func SendFile(conn net.Conn, fName string) {
	defer conn.Close()
	file, err := os.Open(root + fName)
	if err != nil {
		//logger.Log(conn,fName,f)
		return
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println(err)
		return
	}

}
