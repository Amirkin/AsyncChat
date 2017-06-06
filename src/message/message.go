package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type SerializedStruct interface {
	Serialize() []byte
	Deserialize(buffer []byte)
}

type Message struct {
	Nick string
	Text string
	SerializedStruct
}

func (this *Message) Serialize() []byte {
	buf := new(bytes.Buffer)

	WriteStringWithLen(buf, this.Nick)
	WriteStringWithLen(buf, this.Text)

	return buf.Bytes()
}

func (this *Message) Deserialize(buffer []byte) {
	this.Nick, buffer = ReadStringWithLength(buffer)
	this.Text, buffer = ReadStringWithLength(buffer)
}


// TODO: функции что далее, нужно вынести по далее.
func ReadInt(buf []byte) ( value int, buffer []byte) {
	value = int(binary.BigEndian.Uint32(buf[:4]))
	buffer = buf[4:]
	return
}

func ReadString(buf []byte, n int) (value string, buffer []byte) {
	value = string(buf[:n])
	buffer = buf[n:]
	return
}

func ReadStringWithLength(buf []byte) (value string, buffer []byte) {
	n, buf := ReadInt(buf)
	value, buffer = ReadString(buf, n)
	return
}

func WriteInt(buf *bytes.Buffer, value int32) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteString(buf *bytes.Buffer, value string) {
	err := binary.Write(buf, binary.BigEndian, []byte(value))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteStringWithLen(buf *bytes.Buffer, value string) {
	WriteInt(buf, int32(len(value)))
	WriteString(buf, value)
}
