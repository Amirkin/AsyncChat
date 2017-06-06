package client

import (
	"net"
	"fmt"
	"log"
)

type ChatClient struct {
	nickname string
	conn     net.Conn
}

// default text buffer length
const bufferLength = 81920

func NewChatClient(name string) *ChatClient {
	var chatClient ChatClient
	chatClient.nickname = name
	return &chatClient
}

func (this *ChatClient) Connect() {
	var err error
	this.conn, err = net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (this *ChatClient) Write(text string) {
	this.conn.Write([]byte(text))
}

func (this *ChatClient) IsConnect() bool {
	return this.IsConnect()
}

func (this *ChatClient) GetMessage() string {
	log.Println("Called GetMessage() method")

	buffer := make([]byte, bufferLength)

	n, err := this.conn.Read(buffer)
	if err != nil {
		log.Fatalln("Error message reading!")
	}

	this.conn.Write(buffer[:n])
	return string(buffer[:n])
}
