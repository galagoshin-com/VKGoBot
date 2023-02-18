package keyboards

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/requests"
)

type NormalButton struct {
	Row     int
	Column  int
	Text    string
	Payload Payload
	Color   Color
}

func (button NormalButton) getType() ButtonType {
	return normalType
}

func (button NormalButton) getRow() int {
	return button.Row
}

func (button NormalButton) getColumn() int {
	return button.Column
}

func (button NormalButton) getText() string {
	return button.Text
}

func (button NormalButton) getPayload() Payload {
	return button.Payload
}

func (button NormalButton) getColor() Color {
	return button.Color
}

func (button NormalButton) getLink() requests.URL {
	logger.Error(errors.New("Normal button has no link"))
	return ""
}
