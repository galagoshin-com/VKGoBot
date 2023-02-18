package chats

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"net/url"
	"strconv"
)

type CallbackAnswer struct {
	Text string
}

type Callback struct {
	Id      string
	Chat    Chat
	Payload keyboards.Payload
}

func (callback Callback) SendAnswer(answer CallbackAnswer) {
	values := url.Values{
		"v":            {"5.103"},
		"event_id":     {callback.Id},
		"access_token": {tokens.GetToken()},
		"peer_id":      {strconv.Itoa(int(callback.Chat.GetId()))},
		"user_id":      {strconv.Itoa(int(callback.Chat.GetId()))},
		"event_data":   {fmt.Sprintf("{\"type\": \"show_snackbar\", \"text\": \"%s\"}", answer.Text)},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/messages.sendMessageEventAnswer"),
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
