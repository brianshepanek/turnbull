package entity

type message struct {
	*messageStruct
}

type messages struct {
	messagesStruct
}

type Message interface {
	messageInterface
}
type Messages interface {
	messagesInterface
}

func NewMessage() Message {
	return newMessage()
}

func NewMessages() Messages {
	return &messages{}
}
