package message_test

import (
	"testing"
	"message"
	"fmt"
)

func TestMessageSerialize(t *testing.T) {
	msg := message.Message{
		Nick: "Ник",
		Text: "какой то текст",
	}
	buf := msg.Serialize()
	newMsg := message.Message{}
	newMsg.Deserialize(buf)

	fmt.Println(newMsg.Nick + " " + newMsg.Text)
}
