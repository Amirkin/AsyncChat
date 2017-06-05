package client

import (
	"net"
	"fmt"
)

type ChatClient struct {
	nickname string
	conn net.Conn
}

func NewChatClient() *ChatClient {
	var chatClient ChatClient
	return &chatClient
}

func (this * ChatClient) Connect() {
	var err error
	this.conn, err = net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (this * ChatClient) Write(text string) {
	this.conn.Write([]byte(text))
}

func (this *ChatClient) IsConnect() bool {
	return this.IsConnect()
}