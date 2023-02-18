package keyboards

import (
	"github.com/Galagoshin/GoUtils/json"
)

type RemoveKeyboard struct {
	data map[string]any
}

func (keyboard *RemoveKeyboard) Init() {
	data := make(map[string]any)
	keyboard.data = data
}

func (keyboard *RemoveKeyboard) getData() map[string]any {
	return keyboard.data
}

func (keyboard *RemoveKeyboard) setData(data map[string]any) {}

func (keyboard *RemoveKeyboard) GetJson() json.Json {
	return "{\"buttons\":[],\"one_time\":true}"
}
