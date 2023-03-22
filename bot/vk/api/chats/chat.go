package chats

import (
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
)

const absolute_range_bound = 0x7FFFFFFFFFFFFFF

type Chat interface {
	SendMessage(message Message)
	UploadImage(file files.File) attachments.Image
	GetId() uint
}
