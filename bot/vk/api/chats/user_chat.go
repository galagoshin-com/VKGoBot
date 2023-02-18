package chats

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/random"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"net/url"
	"strconv"
)

type UserChat uint

func (chat UserChat) GetId() uint {
	return uint(chat)
}

func (chat UserChat) SendMessage(message Message) {
	values := url.Values{
		"v":            {"5.92"},
		"random_id":    {strconv.Itoa(random.RangeInt(-absolute_range_bound, absolute_range_bound))},
		"access_token": {tokens.GetToken()},
		"peer_id":      {strconv.Itoa(int(chat))},
	}
	if message.Text != "" {
		values.Set("message", message.Text)
	}
	if message.Attachments != nil && len(message.Attachments) > 0 {
		att := ""
		for _, attachment := range message.Attachments {
			att += "," + attachment.BuildString()
		}
		values.Set("attachment", att[1:len(att)])
	}
	if message.Keyboard != nil {
		values.Set("keyboard", string(message.Keyboard.GetJson()))
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/messages.send"),
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
