package keyboards

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"strings"
)

type InlineKeyboard struct {
	data map[string]any
}

func (keyboard *InlineKeyboard) Init() {
	data := make(map[string]any)
	data["inline"] = false
	data["buttons"] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	keyboard.data = data
}

func (keyboard *InlineKeyboard) getData() map[string]any {
	return keyboard.data
}

func (keyboard *InlineKeyboard) setData(data map[string]any) {
	keyboard.data = data
}

func (keyboard *InlineKeyboard) AddButton(button Button) *InlineKeyboard {
	data := keyboard.data
	if data["buttons"].([]any)[button.getRow()] == nil {
		data["buttons"].([]any)[button.getRow()] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()] = make(map[string]any)
	data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"] = make(map[string]any)
	if button.getType() == normalType {
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.getType()
		if len(button.getPayload()) != 0 {
			data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["payload"] = button.getPayload().sign()
		}
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.getText()
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["color"] = button.getColor()
	} else if button.getType() == linkType {
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.getType()
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.getText()
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["link"] = button.getLink()
	} else if button.getType() == callbackType {
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.getType()
		data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.getText()
		if len(button.getPayload()) != 0 {
			data["buttons"].([]any)[button.getRow()].([]any)[button.getColumn()].(map[string]any)["action"].(map[string]any)["payload"] = button.getPayload().sign()
		} else {
			logger.Error(errors.New("Callback button must have payload"))
			return keyboard
		}
	}
	keyboard.setData(data)
	return keyboard
}

func (keyboard *InlineKeyboard) GetJson() json.Json {
	jsn, err := json.Encode(keyboard.data)
	if err != nil {
		logger.Error(err)
	}
	return json.Json(strings.Replace(strings.Replace(string(jsn), ", null", "", -1), ",null", "", -1))
}
