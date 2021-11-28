package ftp

import (
	"fmt"
	"io/ioutil"
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
		fmt.Println(string(input))
		switch cmd[0] {
		case "send":
			fmt.Println("зашел ")
			GetFile(conn, cmd[1], cmd[2])
		case "ls":
			files, err := ioutil.ReadDir("./fileServer")
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, file := range files {
				conn.Write([]byte(file.Name() + "\n"))
			}
		case "get":
			SendFile(conn, cmd[1])

		default:
			fmt.Println("тест")
		}

		break
	}
}
