package handler

import "github.com/Galagoshin/VKGoBot/bot/vk/api/chats"

type Handler interface {
	Init()
	Run()
	RegisterResponser(function func(chats.Chat, chats.OutgoingMessage))
	getEvents(server string, key string, ts string) []byte
	isMessage(response []any, index int) bool
	getObject(response map[string]any) map[string]any
}
