package keyboards

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/requests"
)

type CallbackButton struct {
	Row     int
	Column  int
	Text    string
	Payload Payload
}

func (button CallbackButton) getType() ButtonType {
	return callbackType
}

func (button CallbackButton) getRow() int {
	return button.Row
}

func (button CallbackButton) getColumn() int {
	return button.Column
}

func (button CallbackButton) getText() string {
	return button.Text
}

func (button CallbackButton) getPayload() Payload {
	return button.Payload
}

func (button CallbackButton) getColor() Color {
	logger.Error(errors.New("Callback button has no color"))
	return ""
}

func (button CallbackButton) getLink() requests.URL {
	logger.Error(errors.New("Callback button has no link"))
	return ""
}
