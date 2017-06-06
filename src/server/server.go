package server

import (
	"net"
	"fmt"
	"log"
)

type ChatServer struct {
	listener net.Listener
	clients  []net.Conn
	input    chan []byte
}

func NewChatServer() *ChatServer {
	var chatServer ChatServer
	return &chatServer
}

func (this *ChatServer) accept(conn net.Conn) {
	log.Println("client was connected")
	for {
		buffer := make([]byte, 81920)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		//TODO необходима структура для сообщения, иначе непонятно от кого месседж
		fmt.Println(string(buffer))
		this.input <- buffer
	}
}

func (this *ChatServer) SendAll() {
	for {
		log.Println("SendAll()")
		messageText := <- this.input
		log.Println("messageText = " + string(messageText))
		for _, client := range this.clients {
			client.Write(messageText)
		}
	}
}

func (this *ChatServer) Start() {
	this.input = make(chan []byte)
	var err error
	this.listener, err = net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	go this.SendAll()

	for {
		conn, _ := this.listener.Accept()
		this.clients = append(this.clients, conn)
		go this.accept(conn)
	}
}
