package server

import (
	"net"
	"fmt"
	"log"
)

type ChatServer struct {
	listener net.Listener
}

func NewChatServer() *ChatServer {
	var chatServer ChatServer
	return &chatServer
}

func accept(conn net.Conn) {
	log.Println("client was connected")
	buffer := make([]byte, 81920)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//TODO необходима структура для сообщения, иначе непонятно от кого месседж
	fmt.Println(string(buffer))
}

func (this *ChatServer) Start() {
	var err error
	this.listener, err = net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	for {
		conn, _ := this.listener.Accept()
		go accept(conn)
	}
}
