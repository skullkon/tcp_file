package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/skullkon/tcp_file/cmd/client/fs"
)

func main() {
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:8080")
		defer conn.Close()
		var command string
		fmt.Print("command: ")
		fmt.Scanln(&command)

		switch strings.ToLower(command) {
		case "send":
			fmt.Print("Input filename: ")
			var fileName string
			// var pass string
			fmt.Scanln(&fileName)
			fs.SendFile(conn, fileName, "send")
			break
			// fmt.Println("Sending your " + fileName)

			// fmt.Scanln(&pass)
			// fmt.Println("Sending your " + fileName)

			// case "download":
			// 	fmt.Print("Input filename: ")
			// 	var fileName string
			// 	fmt.Scanln(&fileName)
			// 	fmt.Println("Sending your " + fileName)
			// case "ls":
			// 	fmt.Print("Input filename: ")
			// 	var fileName string
			// 	fmt.Scanln(&fileName)
			// 	fmt.Println("Sending your " + fileName)
		}

	}
}
