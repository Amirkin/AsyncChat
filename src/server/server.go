package server

import (
	"net"
	"log"
	"utils"
	"message"
	"bytes"
)

type User struct {
	nickname string
	conn     net.Conn
}

type ChatServer struct {
	listener net.Listener
	users    []User
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

		var msg string
		cmd, buf := command.GetCommand(buffer[:n])
		switch cmd {
		case command.NICK:
			{
				name, _ := message.ReadStringWithLength(buf)
				log.Println(" В чат вошёл: " + name)
				msg = " В чат вошёл: " + name
			}
		case command.MSG:
			{
				msg, _ = message.ReadStringWithLength(buf)
			}
		}

		this.input <- []byte(msg)
	}
}

func (this *ChatServer) SendAll() {
	for {
		log.Println("SendAll()")
		messageText := <-this.input
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
