package events

import "github.com/Galagoshin/GoUtils/events"

const (
	StopApplicationEvent = events.EventName("StopApplicationEvent")
	HotReloadEvent       = events.EventName("HotReloadEvent")
	EnablePluginEvent    = events.EventName("EnablePluginEvent")
	DisablePluginEvent   = events.EventName("DisablePluginEvent")
	StopBotEvent         = events.EventName("StopBotEvent")
	StartBotEvent        = events.EventName("StartBotEvent")

	MessageCallbackEvent = events.EventName("MessageCallbackEvent")
)
