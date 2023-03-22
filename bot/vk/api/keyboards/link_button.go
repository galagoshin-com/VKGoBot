package keyboards

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/requests"
)

type LinkButton struct {
	Row    int
	Column int
	Text   string
	Link   requests.URL
}

func (button LinkButton) GetType() ButtonType {
	return LinkType
}

func (button LinkButton) GetPayload() Payload {
	logger.Error(errors.New("Link button has no payload"))
	return make(map[string]any)
}

func (button LinkButton) GetColor() Color {
	logger.Error(errors.New("Link button has no color"))
	return ""
}

func (button LinkButton) GetLink() requests.URL {
	return button.Link
}

func (button LinkButton) GetText() string {
	return button.Text
}

func (button LinkButton) GetRow() int {
	return button.Row
}

func (button LinkButton) GetColumn() int {
	return button.Column
}
