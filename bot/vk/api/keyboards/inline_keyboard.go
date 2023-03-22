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
	data["inline"] = true
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
	if data["buttons"].([]any)[button.GetRow()] == nil {
		data["buttons"].([]any)[button.GetRow()] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()] = make(map[string]any)
	data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"] = make(map[string]any)
	if button.GetType() == NormalType {
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
		if len(button.GetPayload()) != 0 {
			data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["payload"] = button.GetPayload().Sign()
		}
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["color"] = button.GetColor()
	} else if button.GetType() == LinkType {
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["link"] = button.GetLink()
	} else if button.GetType() == CallbackType {
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
		data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
		if len(button.GetPayload()) != 0 {
			data["buttons"].([]any)[button.GetRow()].([]any)[button.GetColumn()].(map[string]any)["action"].(map[string]any)["payload"] = button.GetPayload().Sign()
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
