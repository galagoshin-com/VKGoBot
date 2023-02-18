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

func (button LinkButton) getType() ButtonType {
	return linkType
}

func (button LinkButton) getPayload() Payload {
	logger.Error(errors.New("Link button has no payload"))
	return make(map[string]any)
}

func (button LinkButton) getColor() Color {
	logger.Error(errors.New("Link button has no color"))
	return ""
}

func (button LinkButton) getLink() requests.URL {
	return button.Link
}

func (button LinkButton) getText() string {
	return button.Text
}

func (button LinkButton) getRow() int {
	return button.Row
}

func (button LinkButton) getColumn() int {
	return button.Column
}
