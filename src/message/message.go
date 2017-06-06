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

	writeStringWithLen(buf, this.Nick)
	writeStringWithLen(buf, this.Text)

	return buf.Bytes()
}

func (this *Message) Deserialize(buffer []byte) {
	this.Nick, buffer = readStringWithLength(buffer)
	this.Text, buffer = readStringWithLength(buffer)
}

func readInt(buf []byte) ( value int, buffer []byte) {
	value = int(binary.BigEndian.Uint32(buf[:4]))
	buffer = buf[4:]
	return
}

func readString(buf []byte, n int) (value string, buffer []byte) {
	value = string(buf[:n])
	buffer = buf[n:]
	return
}

func readStringWithLength(buf []byte) (value string, buffer []byte) {
	n, buf := readInt(buf)
	value, buffer = readString(buf, n)
	return
}

func writeInt(buf *bytes.Buffer, value int32) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func writeString(buf *bytes.Buffer, value string) {
	err := binary.Write(buf, binary.BigEndian, []byte(value))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func writeStringWithLen(buf *bytes.Buffer, value string) {
	writeInt(buf, int32(len(value)))
	writeString(buf, value)
}
