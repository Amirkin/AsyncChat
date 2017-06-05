package main

import (
	"os"
	"server"
	"client"
)

func main() {
	if len(os.Args) == 2 {

		switch os.Args[1] {
		case "server":
			{
				server := server.NewChatServer()
				server.Start()
			}
		case "client":
			{
				client := client.NewChatClient()
				client.Connect()
				client.Write("какой то текст")
			}
		}

	}
}
