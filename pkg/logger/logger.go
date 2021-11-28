package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

var (
	file *os.File
	err  error
)

func init() {
	file, err = os.Create("./log.txt")
	if err != nil {
		logrus.Error(err)
	}
}

// Log status == 0 отправка, status == 1 получение
func Log(conn net.Conn, fName string, fSize string, status bool, err error) {
	log := fmt.Sprintf("%s  %s  %s  %t  %v", time.Now().String(), fName, fSize, status, err.Error())
	file.Write([]byte(log))
}
