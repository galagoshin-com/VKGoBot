package chats

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/templates"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
)

type OutgoingMessage struct {
	Text        string
	User        users.User
	Payload     keyboards.Payload
	Attachments []attachments.Attachment
	Client      Client
}

type Message struct {
	Text        string
	Template    templates.Template
	Keyboard    keyboards.Keyboard
	Attachments []attachments.Attachment
}
