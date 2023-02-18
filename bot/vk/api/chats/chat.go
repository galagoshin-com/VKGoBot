package chats

const absolute_range_bound = 0x7FFFFFFFFFFFFFF

type Chat interface {
	SendMessage(message Message)
	GetId() uint
}
