package client

import (
	"net"
	"fmt"
	"log"
	"utils"
)


type ChatClient struct {
	nickname string
	conn     net.Conn
}

// default text buffer length
const bufferLength = 81920

func NewChatClient(nickname string) *ChatClient {
	var chatClient ChatClient
	chatClient.nickname = nickname
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

func (this *ChatClient) SendNick() {
	command.SendCommandString(this.conn, command.NICK, this.nickname)
}

func (this *ChatClient) IsConnect() bool {
	return this.IsConnect()
}

func (this *ChatClient) GetByteMessage() []byte {
	log.Println("Called GetByteMessage() method")

	buffer := make([]byte, bufferLength)

	n, err := this.conn.Read(buffer)
	if err != nil {
		log.Fatalln("Error message reading!")
	}

	return buffer[:n]
}

func (this *ChatClient) GetMessage() string {
	log.Println("Called GetMessage() method")

	buffer := make([]byte, bufferLength)

	n, err := this.conn.Read(buffer)
	if err != nil {
		log.Fatalln("Error message reading!")
	}

	return string(buffer[:n])
}

func (this *ChatClient) GetNickName() string {
	return this.nickname
}

func (this *ChatClient) SendMessage(msg string) {
	command.SendCommandString(this.conn, command.MSG, msg)
}
