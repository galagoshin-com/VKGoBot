package comments

import (
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
)

type Commentator int

type Comment struct {
	Commentator   Commentator
	CommentObject attachments.Attachment
	Date          int64
	Text          string
}
