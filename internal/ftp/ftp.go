package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if err != nil {
			return
		}

		str := string(input[0:n])
		cmd := strings.Split(str, "\n")
		switch cmd[0] {
		case "send":
			fmt.Println("зашел ")
			GetFile(conn, cmd[1])
		case "ls":
			files, err := ioutil.ReadDir("./fileServer")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				conn.Write([]byte(file.Name() + "\n"))
			}

		default:
			fmt.Println("тест")
		}

		break
	}
}
