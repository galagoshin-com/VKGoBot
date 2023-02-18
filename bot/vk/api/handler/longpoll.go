package handler

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/chats"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/groups"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/keyboards"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/users"
	"net/url"
	"strconv"
)

type LongPoll struct {
	server    string
	key       string
	ts        string
	events    []byte
	onMessage func(chats.Chat, chats.OutgoingMessage)
}

var Group groups.Group

func (longpoll *LongPoll) Init() {
	logger.Print("Initializing LongPoll server...")
	values := url.Values{
		"v":            {"5.113"},
		"access_token": {tokens.GetToken()},
		"group_id":     {strconv.Itoa(int(Group))},
	}
	request := requests.Request{
		Method: requests.POST,
		Url:    requests.URL("https://api.vk.com/method/groups.getLongPollServer"),
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
	if response_json["response"] == nil {
		logger.Error(errors.New("Error while init longpoll server"))
	}
	md := response_json["response"].(map[string]any)
	ts := md["ts"].(string)
	key := md["key"].(string)
	server := md["server"].(string)
	longpoll.server = server
	longpoll.key = key
	longpoll.ts = ts
	longpoll.events = []byte{}
	logger.Print("LongPoll server initialized!")
}

func (longpoll *LongPoll) RegisterResponser(function func(chats.Chat, chats.OutgoingMessage)) {
	longpoll.onMessage = function
}

func (longpoll *LongPoll) Run() {
	logger.Print("LongPoll server started!")
	first := true
	for true {
		var ts string
		if first {
			ts = longpoll.ts
		} else {
			tsjson, err := json.Decode(json.Json(longpoll.events))
			if err != nil {
				logger.Error(err)
			}
			ts = tsjson["ts"].(any).(string)
		}
		longpoll.events = longpoll.getEvents(longpoll.server, longpoll.key, ts)
		updatesjsn, err := json.Decode(json.Json(longpoll.events))
		if err != nil {
			logger.Error(err)
		}
		if updatesjsn["updates"] != nil {
			msgs := updatesjsn["updates"].([]any)
			for msg := range msgs {
				if longpoll.isMessage(msgs, msg) {
					data := longpoll.getObject(msgs[msg].(map[string]any))["message"].(map[string]any)
					client_info := longpoll.getObject(msgs[msg].(map[string]any))["client_info"].(map[string]any)
					res := []string{}
					for _, button := range client_info["button_actions"].([]any) {
						res = append(res, button.(string))
					}
					client := chats.Client{
						ButtonActions:  res,
						Keyboard:       client_info["keyboard"].(bool),
						InlineKeyboard: client_info["inline_keyboard"].(bool),
						Carousel:       client_info["carousel"].(bool),
					}
					payload := keyboards.Payload{}
					if data["payload"] != nil {
						payload, err = json.Decode(json.Json(data["payload"].(string)))
						if payload["sign"] != nil {
							if err != nil {
								logger.Error(err)
							}
							if !payload.Verify() {
								logger.Warning("Incorrect payload sign!")
								continue
							}
						}
					}
					if int(data["from_id"].(float64)) > 0 {
						outgoing := chats.OutgoingMessage{
							Text:    data["text"].(string),
							User:    users.User(data["from_id"].(float64)),
							Payload: payload,
							Client:  client,
						}
						if longpoll.onMessage != nil {
							if data["peer_id"] == data["from_id"] {
								logger.Print(fmt.Sprintf("user%d: %s", outgoing.User, outgoing.Text))
								logger.Debug(1, false, fmt.Sprintf("Message payload: %v", outgoing.Payload))
								longpoll.onMessage(chats.UserChat(outgoing.User), outgoing)
							} else {
								logger.Print(fmt.Sprintf("conv%d: %s", outgoing.User, outgoing.Text))
								logger.Debug(1, false, fmt.Sprintf("Message payload: %v", outgoing.Payload))
								longpoll.onMessage(chats.ConversationChat(data["peer_id"].(float64)), outgoing)
							}
						} else {
							logger.Warning("No responser found! Ignoring messages...")
						}
					}
				} else if longpoll.isMessageEvent(msgs, msg) {
					data := longpoll.getObject(msgs[msg].(map[string]any))
					payload := keyboards.Payload{}
					if data["payload"] != nil {
						payload = data["payload"].(map[string]any)
						if payload["sign"] != nil {
							if err != nil {
								logger.Error(err)
							}
							if !payload.Verify() {
								logger.Warning("Incorrect payload sign!")
								continue
							}
						}
					}
					if int(data["user_id"].(float64)) > 0 {
						if data["peer_id"] == data["user_id"] {
							callback := chats.Callback{
								Id:      data["event_id"].(string),
								Chat:    chats.UserChat(data["user_id"].(float64)),
								Payload: payload,
							}
							logger.Debug(1, false, fmt.Sprintf("Event payload: %v", payload))
							events.CallAllEvents(events2.MessageCallbackEvent, callback)
						} else {
							callback := chats.Callback{
								Id:      data["event_id"].(string),
								Chat:    chats.ConversationChat(data["peer_id"].(float64)),
								Payload: payload,
							}
							logger.Debug(1, false, fmt.Sprintf("Event payload: %v", payload))
							events.CallAllEvents(events2.MessageCallbackEvent, callback)
						}
					}
				}
			}
			if first {
				first = false
			}
		} else {
			longpoll.Init()
			first = true
		}
	}
}

func (longpoll *LongPoll) getEvents(server, key, ts string) []byte {
	getEventsLocaly := func(server, key, ts string) ([]byte, error) {
		request := requests.Request{
			Method: requests.GET,
			Url:    requests.URL(server + "?act=a_check&key=" + key + "&ts=" + ts + "&wait=25"),
		}
		response, err := request.Send()
		if err != nil {
			logger.Error(err)
			return []byte{}, err
		}
		response_json, err := json.Decode(json.Json(response.Text()))
		if err != nil {
			logger.Error(err)
			return []byte{}, err
		}
		if response_json["error"] != nil {
			logger.Error(errors.New(response_json["error"].(map[string]any)["error_msg"].(string)))
			return []byte{}, errors.New(response_json["error"].(map[string]any)["error_msg"].(string))
		}
		return response.Body, nil
	}
	result, err := getEventsLocaly(server, key, ts)
	if err != nil {
		longpoll.Init()
		result, err = getEventsLocaly(longpoll.server, longpoll.key, longpoll.ts)
		if err != nil {
			logger.Error(err)
		} else {
			return result
		}
	} else {
		return result
	}
	return []byte{}
}

func (longpoll *LongPoll) isMessage(response []any, index int) bool {
	if len(response) != 0 {
		md := response[index].(map[string]any)
		return md["type"] == "message_new"
	} else {
		return false
	}
}

func (longpoll *LongPoll) isMessageEvent(response []any, index int) bool {
	if len(response) != 0 {
		md := response[index].(map[string]any)
		return md["type"] == "message_event"
	} else {
		return false
	}
}

func (longpoll *LongPoll) getObject(response map[string]any) map[string]any {
	return response["object"].(map[string]any)
}
