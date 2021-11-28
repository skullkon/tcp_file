package ftp

import (
	"fmt"
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
		default:
			fmt.Println("тест")
		}

		break
	}
}
