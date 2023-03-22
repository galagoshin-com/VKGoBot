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

func (button NormalButton) GetType() ButtonType {
	return NormalType
}

func (button NormalButton) GetRow() int {
	return button.Row
}

func (button NormalButton) GetColumn() int {
	return button.Column
}

func (button NormalButton) GetText() string {
	return button.Text
}

func (button NormalButton) GetPayload() Payload {
	return button.Payload
}

func (button NormalButton) GetColor() Color {
	return button.Color
}

func (button NormalButton) GetLink() requests.URL {
	logger.Error(errors.New("Normal button has no link"))
	return ""
}
