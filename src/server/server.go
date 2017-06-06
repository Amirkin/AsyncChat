package server

import (
	"net"
	"log"
	"message"
	"utils"
	"bytes"
)

type DataCmd struct {
	cmd  int
	data []byte
}

func (this DataCmd) Serialize() []byte {
	buf := new(bytes.Buffer)
	message.WriteInt(buf, int32(this.cmd))
	buf.Write(this.data)
	return buf.Bytes()
}

type User struct {
	nickname string
	conn     net.Conn
}

type ChatServer struct {
	listener net.Listener
	users    []User
	input    chan DataCmd
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

		cmd, buf := command.GetCommand(buffer[:n])
		switch cmd {
			case command.NICK:
			{
				data := DataCmd{}
				user.nickname, _ = message.ReadStringWithLength(buf)

				log.Println(" В чат вошёл: " + user.nickname)
				data.cmd = command.NEWUSER
				data.data = buf
				this.input <- data
			}
			case command.MSG:
			{
				data := DataCmd{}
				data.cmd = command.MSG
				msg := message.Message{}
				msg.Nick = user.nickname
				msg.Text = string(buf[4:])
				data.data = msg.Serialize()

				this.input <- data
			}
		}
	}
}

func (this *ChatServer) SendAll() {
	for {
		log.Println("SendAll()")
		data := <- this.input

		for _, user := range this.users {
			user.conn.Write(data.Serialize())
		}
	}
}

func (this *ChatServer) Start() {
	this.input = make(chan DataCmd)
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
