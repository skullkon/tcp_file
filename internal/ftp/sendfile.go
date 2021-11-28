package ftp

//
//import (
//	"github.com/google/uuid"
//	"io"
//	"log"
//	"net"
//	"os"
//	"strings"
//)
//
//const (
//	root = "./fileServer/"
//)
//
//func SendFile(conn net.Conn, fName string) {
//
//	file, err := os.Open(root + uuid.New().String())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer file.Close()
//
//	_, err = io.Copy(file, conn)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	conn.Close()
//}
