package chats

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/random"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type ConversationChat uint

func (chat ConversationChat) GetId() uint {
	return uint(chat)
}

func (chat ConversationChat) SendMessage(message Message) {
	if chat < 2000000000 {
		logger.Error(errors.New(fmt.Sprintf("%d is not conversation id", chat)))
		return
	}
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
	if message.Template != nil {
		values.Set("template", string(message.Template.GetJson()))
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

func (chat ConversationChat) UploadImage(file files.File) attachments.Image {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var client *http.Client
	var remoteURL string
	client = &http.Client{Transport: tr}
	request := requests.Request{
		Url:    "https://api.vk.com/method/photos.getMessagesUploadServer",
		Method: requests.POST,
		Data: url.Values{
			"access_token": {tokens.GetToken()},
			"v":            {"5.131"},
			"peer_id":      {strconv.Itoa(int(chat))},
		},
	}
	response, err := request.Send()
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}
	json_response, err := json.Decode(json.Json(response.Text()))
	remoteURL = json_response["response"].(map[string]interface{})["upload_url"].(string)
	r, err := os.Open(file.Path)
	if err != nil {
		logger.Error(err)
	}
	values := map[string]io.Reader{
		"photo": r,
	}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			fw, _ = w.CreateFormFile(key, x.Name())
		} else {
			fw, _ = w.CreateFormField(key)
		}
		io.Copy(fw, r)

	}
	w.Close()
	req, err := http.NewRequest("POST", remoteURL, &b)
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	bts, _ := ioutil.ReadAll(res.Body)

	photo_jsn, err := json.Decode(json.Json(bts))
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}

	photo := photo_jsn["photo"].(string)
	server := strconv.Itoa(int(photo_jsn["server"].(float64)))
	hash := photo_jsn["hash"].(string)

	request = requests.Request{
		Url:    "https://api.vk.com/method/photos.saveMessagesPhoto",
		Method: requests.POST,
		Data: url.Values{
			"access_token": {tokens.GetToken()},
			"v":            {"5.131"},
			"photo":        {photo},
			"server":       {server},
			"hash":         {hash},
		},
	}
	result, err := request.Send()
	if err != nil {
		logger.Error(err)
		return attachments.Image{}
	}

	json_res, err := json.Decode(json.Json(result.Text()))

	id := uint(json_res["response"].([]any)[0].(map[string]any)["id"].(float64))
	owner_id := int(json_res["response"].([]any)[0].(map[string]any)["owner_id"].(float64))
	access_key := json_res["response"].([]any)[0].(map[string]any)["access_key"].(string)
	return attachments.Image{
		OwnerId:   owner_id,
		Id:        id,
		AccessKey: access_key,
	}
}
