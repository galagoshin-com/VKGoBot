package keyboards

import "github.com/Galagoshin/GoUtils/requests"

type ButtonType string

const (
	NormalType   = ButtonType("text")
	LinkType     = ButtonType("open_link")
	CallbackType = ButtonType("callback")
)

type Button interface {
	GetType() ButtonType
	GetRow() int
	GetColumn() int
	GetPayload() Payload
	GetText() string
	GetLink() requests.URL
	GetColor() Color
}
