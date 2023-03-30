package vk

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/configs"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/files"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/groups"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/handler"
	tokens2 "github.com/Galagoshin/VKGoBot/bot/vk/api/tokens"
	"os"
	"strings"
)

var Config = &configs.Config{Name: "bot"}

var (
	groupId int = 0
)

var vkhandler handler.Handler

func Init() {
	Config.Init(map[string]any{
		"group-id":           123456789,
		"group-tokens-list":  "group_tokens.list",
		"user-tokens-list":   "user_tokens.list",
		"widget-tokens-list": "widget_tokens.list",
	}, 1)

	group_id, group_id_error := Config.Get("group-id")
	if !group_id_error {
		logger.Panic(errors.New("\"group-id\" is not defined in the bot config."))
	} else {
		groupId = group_id.(int)
	}
	group_tokens_list, group_tokens_list_error := Config.Get("group-tokens-list")
	if !group_tokens_list_error {
		logger.Panic(errors.New("\"group-tokens-list\" is not defined in the bot config."))
	} else {
		file := files.File{Path: group_tokens_list.(string)}
		if !file.Exists() {
			err := file.Create()
			if err != nil {
				logger.Panic(err)
			}
			err = file.WriteString("\n")
			if err != nil {
				logger.Panic(err)
			}
		}
		err := file.Open(os.O_RDWR)
		if err != nil {
			logger.Panic(err)
		}
		tokens := file.ReadString()
		tokens2.GroupTokensList = strings.Split(tokens, "\n")
	}
	user_tokens_list, user_tokens_list_error := Config.Get("user-tokens-list")
	if !user_tokens_list_error {
		logger.Panic(errors.New("\"user-tokens-list\" is not defined in the bot config."))
	} else {
		file := files.File{Path: user_tokens_list.(string)}
		if !file.Exists() {
			err := file.Create()
			if err != nil {
				logger.Panic(err)
			}
			err = file.WriteString("\n")
			if err != nil {
				logger.Panic(err)
			}
		}
		err := file.Open(os.O_RDWR)
		if err != nil {
			logger.Panic(err)
		}
		tokens := file.ReadString()
		tokens2.UserTokensList = strings.Split(tokens, "\n")
	}
	widget_tokens_list, widget_tokens_list_error := Config.Get("widget-tokens-list")
	if !widget_tokens_list_error {
		logger.Panic(errors.New("\"widget-tokens-list\" is not defined in the bot config."))
	} else {
		file := files.File{Path: widget_tokens_list.(string)}
		if !file.Exists() {
			err := file.Create()
			if err != nil {
				logger.Panic(err)
			}
			err = file.WriteString("\n")
			if err != nil {
				logger.Panic(err)
			}
		}
		err := file.Open(os.O_RDWR)
		if err != nil {
			logger.Panic(err)
		}
		tokens := file.ReadString()
		tokens2.WidgetTokensList = strings.Split(tokens, "\n")
	}
	vkhandler = &handler.LongPoll{}
	handler.Group = groups.Group(groupId)
	vkhandler.Init()
}

func Stop() {
	events.CallAllEvents(events2.StopBotEvent)
}

func Run() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r.(error))
			vkhandler.Init()
			Run()
		}
	}()
	events.CallAllEvents(events2.StartBotEvent)
	vkhandler.Run()
}

func GetHandler() handler.Handler {
	return vkhandler
}
