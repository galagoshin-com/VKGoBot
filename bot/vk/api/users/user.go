package users

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"net/url"
	"strconv"
)

type User uint

func ExistsUser(user User) bool {
	values := url.Values{
		"v":            {"5.103"},
		"access_token": {tokens.GetToken()},
		"user_ids":     {strconv.Itoa(int(user))},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/users.get"),
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
	if response_json["response"] != nil {
		return response_json["response"].([]interface{})[0] == nil
	} else {
		return false
	}
}
