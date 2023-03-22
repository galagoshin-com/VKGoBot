package templates

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"strings"
)

type CarouselAction string

const (
	OpenLink  = CarouselAction("open_link")
	OpenPhoto = CarouselAction("open_photo")
)

type Carousel struct {
	data map[string]any
}

func (keyboard *Carousel) Init() {
	data := make(map[string]any)
	data["type"] = "carousel"
	data["elements"] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	keyboard.data = data
}

func (keyboard *Carousel) getData() map[string]any {
	return keyboard.data
}

func (keyboard *Carousel) setData(data map[string]any) {
	keyboard.data = data
}

func (keyboard *Carousel) AddElement(index int, title, description string, image attachments.Image, buttons []keyboards.Button) *Carousel {
	data := keyboard.data
	if data["elements"].([]any)[index] == nil {
		data["elements"].([]any)[index] = make(map[string]any)
	}
	data["elements"].([]any)[index].(map[string]any)["buttons"] = []any{nil, nil, nil}
	data["elements"].([]any)[index].(map[string]any)["title"] = title
	data["elements"].([]any)[index].(map[string]any)["description"] = description
	if image.Id != 0 {
		data["elements"].([]any)[index].(map[string]any)["photo_id"] = fmt.Sprintf("%d_%d", image.OwnerId, image.Id)
	}
	//data["elements"].([]any)[index].(map[string]any)["action"] = action
	for _, button := range buttons {
		data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()] = make(map[string]any)
		data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"] = make(map[string]any)
		if button.GetType() == keyboards.NormalType {
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
			if len(button.GetPayload()) != 0 {
				data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["payload"] = button.GetPayload().Sign()
			}
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["color"] = button.GetColor()
		} else if button.GetType() == keyboards.LinkType {
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["link"] = button.GetLink()
		} else if button.GetType() == keyboards.CallbackType {
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["type"] = button.GetType()
			data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["label"] = button.GetText()
			if len(button.GetPayload()) != 0 {
				data["elements"].([]any)[index].(map[string]any)["buttons"].([]any)[button.GetRow()].(map[string]any)["action"].(map[string]any)["payload"] = button.GetPayload().Sign()
			} else {
				logger.Error(errors.New("Callback button must have payload"))
				return keyboard
			}
		}
	}
	keyboard.setData(data)
	return keyboard
}

func (keyboard *Carousel) GetJson() json.Json {
	jsn, err := json.Encode(keyboard.data)
	if err != nil {
		logger.Error(err)
	}
	return json.Json(strings.Replace(strings.Replace(string(jsn), ", null", "", -1), ",null", "", -1))
}
