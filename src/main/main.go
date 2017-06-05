package main

import (
	"os"
	"server"
	"client"
	"fmt"
	"sync"
)

func loopGetMessage(client *client.ChatClient) {
	client.Write("какой то текст")
	for {
		msg := client.GetMessage()
		fmt.Println(msg)
	}
}

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
				client := client.NewChatClient("mainClient")
				client.Connect()
				var wg sync.WaitGroup
				wg.Add(1)
				go loopGetMessage(client)
				wg.Wait()
			}
		}

	}
}
