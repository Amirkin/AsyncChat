package server

import (
	"net"
	"fmt"
	"log"
	"strings"
)

type User struct {
	nickname string
	conn net.Conn
}

type ChatServer struct {
	listener net.Listener
	users []User
	input    chan []byte
}

func NewChatServer() *ChatServer {
	var chatServer ChatServer
	return &chatServer
}

func (this *ChatServer) accept(user User) {
	log.Println("client was connected")
	buffer := make([]byte, 81920)
	for {
		n, err := user.conn.Read(buffer)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		//TODO необходима структура для сообщения, иначе непонятно от кого месседж
		msg := string(buffer[:n])
		fmt.Println(msg)

		if strings.HasPrefix(msg, "NICK ") {
			msg = strings.TrimPrefix(msg, "NICK ")
			user.nickname = msg
			msg = "NEWUSER " + user.nickname
		}

		if strings.HasPrefix(msg, "MSG ") {
			msg = "MSG " + user.nickname + ": " + strings.TrimPrefix(msg, "MSG ")
		}


		this.input <- []byte(msg)
	}
}

func (this *ChatServer) SendAll() {
	for {
		log.Println("SendAll()")
		messageText := <- this.input
		log.Println("messageText = " + string(messageText))

		for _, user := range this.users {
			user.conn.Write(messageText)
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
		var user User
		user.conn = conn
		this.users = append(this.users, user)

		go this.accept(user)
	}
}
