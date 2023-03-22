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

func (button CallbackButton) GetType() ButtonType {
	return CallbackType
}

func (button CallbackButton) GetRow() int {
	return button.Row
}

func (button CallbackButton) GetColumn() int {
	return button.Column
}

func (button CallbackButton) GetText() string {
	return button.Text
}

func (button CallbackButton) GetPayload() Payload {
	return button.Payload
}

func (button CallbackButton) GetColor() Color {
	logger.Error(errors.New("Callback button has no color"))
	return ""
}

func (button CallbackButton) GetLink() requests.URL {
	logger.Error(errors.New("Callback button has no link"))
	return ""
}
