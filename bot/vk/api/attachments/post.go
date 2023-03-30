package attachments

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"net/url"
)

type Post struct {
	OwnerId   int
	Id        uint
	AccessKey string
	Text      string
	Date      int64
}

func (post *Post) GetType() AttachmentType {
	return PostType
}

func (post *Post) GetOwnerId() int {
	return post.OwnerId
}

func (post *Post) GetId() uint {
	return post.Id
}

func (post *Post) GetAccessKey() string {
	return post.AccessKey
}

func (post *Post) BuildString() string {
	if post.AccessKey == "" {
		return fmt.Sprintf("%s%d_%d", post.GetType(), post.OwnerId, post.Id)
	} else {
		return fmt.Sprintf("%s%d_%d_%s", post.GetType(), post.OwnerId, post.Id, post.AccessKey)
	}
}

func (post *Post) Init() {
	values := url.Values{
		"v":            {"5.103"},
		"access_token": {tokens.GetUserToken()},
		"posts":        {string([]rune(post.BuildString())[4:])},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/wall.getById"),
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
	post.Text = response_json["response"].([]any)[0].(map[string]any)["text"].(string)
	post.Date = int64(response_json["response"].([]any)[0].(map[string]any)["date"].(float64))
}
