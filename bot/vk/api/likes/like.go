package likes

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
)

type Like struct {
	Liker       users.User
	LikedObject attachments.Attachment
}
