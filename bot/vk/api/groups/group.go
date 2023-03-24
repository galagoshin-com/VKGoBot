package groups

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/widgets"
	"net/url"
)

type Group uint

func (group Group) SetStatus(text string) {
	values := url.Values{
		"v":            {"5.103"},
		"access_token": {tokens.GetUserToken()},
		"group_id":     {fmt.Sprintf("%d", group)},
		"text":         {text},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/status.set"),
		Data:   values,
	}
	response, err := request.Send()
	if err != nil {
		logger.Error(err)
	}
	response_json, err := json.Decode(json.Json(response.Text()))
	if err != nil {
		logger.Error(err)
	}
	if response_json["error"] != nil {
		logger.Error(errors.New(response_json["error"].(map[string]any)["error_msg"].(string)))

	}
}

func (group Group) SetWidget(widjet widgets.Widget) {
	code := url.QueryEscape("return " + string(widjet.GetJson()) + ";")
	values := url.Values{
		"v":            {"5.103"},
		"access_token": {tokens.GetWidgetToken()},
		"code":         {code},
		"type":         {string(widjet.GetType())},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/appWidgets.update"),
		Data:   values,
	}
	response, err := request.Send()
	if err != nil {
		logger.Error(err)
	}
	response_json, err := json.Decode(json.Json(response.Text()))
	if err != nil {
		logger.Error(err)
	}
	if response_json["error"] != nil {
		logger.Error(errors.New(response_json["error"].(map[string]any)["error_msg"].(string)))

	}
}
