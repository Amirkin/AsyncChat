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
	"utils"
	"message"
)

func protocol(client *client.ChatClient) {
	client.SendNick()
}

func loopGetMessage(client *client.ChatClient) {
	for {
		cmd, buf := command.GetCommand(client.GetByteMessage())
		switch cmd {
		case command.MSG:
			{
				msg := message.Message{}
				msg.Deserialize(buf)
				fmt.Println(msg.Nick + ": " + msg.Text)
			}
		case command.NEWUSER:
			{
				name, _ := message.ReadStringWithLength(buf)
				fmt.Println(" В чат вошёл: " + name)
			}
		}
	}
}

func inputMessage(client *client.ChatClient) {
	in := bufio.NewReader(os.Stdin)
	for {
		line, _ := in.ReadString('\n')
		line = strings.Trim(line, "\n")
		client.SendMessage(line)
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
