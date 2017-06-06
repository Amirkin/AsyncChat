package main

import (
	"os"
	"server"
	"client"
	"fmt"
	"bufio"
	"strings"
)

func loopGetMessage(client *client.ChatClient) {
	client.Write("В чат вошёл: " + client.GetNickName())
	for {
		msg := client.GetMessage()
		fmt.Println(msg)
	}
}

func inputMessage(client *client.ChatClient) {
	in := bufio.NewReader(os.Stdin)
	for {
		line, _ := in.ReadString('\n')
		line = strings.Trim(line, "\n")
		client.Write(line)
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
				client := client.NewChatClient(enterNickName())
				client.Connect()
				go loopGetMessage(client)
				inputMessage(client)
			}
		}

	}
}
func enterNickName() string {
	var nickname string
	fmt.Println("Как вас зовут?")
	fmt.Scanf("%s", &nickname)
	return nickname
}
