package main

import (
	"os"
	"server"
	"client"
	"fmt"
	"bufio"
	"strings"
	"log"
	"io"
)

func protocol(client *client.ChatClient) {
	client.SendNick()
}

func loopGetMessage(client *client.ChatClient) {
	for {
		msg := client.GetMessage()
		if strings.HasPrefix(msg, "MSG ") {
			msg = strings.TrimPrefix(msg, "MSG ")
			fmt.Println(msg)
		}
		if strings.HasPrefix(msg, "NEWUSER ") {
			msg = strings.TrimPrefix(msg, "NEWUSER ")
			fmt.Println(" В чат вошёл: " + msg)
		}
	}
}

func inputMessage(client *client.ChatClient) {
	in := bufio.NewReader(os.Stdin)
	for {
		line, _ := in.ReadString('\n')
		line = strings.Trim(line, "\n")
		client.SendMessage(line)
		client.Write(line)
	}
}


func main() {
	logFile, _ := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND, 0)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

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
				protocol(client)
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
