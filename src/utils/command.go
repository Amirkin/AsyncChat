package command

import (
	"message"
	"bytes"
	"net"
)

const (
	NEWUSER = 1
	MSG = 2
	NICK = 3
)

func GetCommand(buf []byte) (cmd int, buffer []byte) {
	cmd, buffer = message.ReadInt(buf)
	return
}

func SendCommand(conn net.Conn, cmd int, object message.SerializedStruct) {
	buf := new(bytes.Buffer)
	message.WriteInt(buf, int32(cmd))
	buf.Write(object.Serialize())
	conn.Write(buf.Bytes())
}

func SendCommandString(conn net.Conn, cmd int, value string) {
	buf := new(bytes.Buffer)
	message.WriteInt(buf, int32(cmd))
	message.WriteStringWithLen(buf, value)
	conn.Write(buf.Bytes())
}
