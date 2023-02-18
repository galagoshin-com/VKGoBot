package keyboards

import "github.com/Galagoshin/GoUtils/requests"

type ButtonType string

const (
	normalType   = ButtonType("text")
	linkType     = ButtonType("open_link")
	callbackType = ButtonType("callback")
)

type Button interface {
	getType() ButtonType
	getRow() int
	getColumn() int
	getPayload() Payload
	getText() string
	getLink() requests.URL
	getColor() Color
}
