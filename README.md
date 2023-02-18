# VKGoBot
Open-source, high-performance framework for the creating VK bots.

### Contents:
- [Getting started](#getting-started)
    - [Install](#install)
    - [Example usage](#example-usage)
    - [Bot configuration](#bot-configuration)
    - [Framework configuration](#framework-configuration)
    - [Developer mode](#developer-mode)
    - [Live reloads](#live-reloads)
    - [Hot reload](#hot-reload)
- [Chats](#chats)
    - [Example responser](#example-responser)
    - [Attachments](#attachments)
    - [Keyboards](#keyboards)
      - [Init keyboard](#init-keyboard)
      - [Buttons](#buttons)
          - [Add buttons](#add-buttons)
          - [Normal button](#normal-button)
          - [Callback button](#normal-button)
          - [Link button](#link-button)
      - [Static keyboard](#static-keyboard)
      - [Inline keyboard](#inline-keyboard)
      - [One time keyboard](#one-time-keyboard)
      - [Remove keyboard](#remove-keyboard)
    - [Messages](#messages)
- [Plugins](#plugins)
  - [Example plugin](#example-plugin)

## Getting started

### Install

`go get -t github.com/Galagoshin/VKGoBot`

### Example usage

```go
package main

import (
  "github.com/Galagoshin/VKGoBot/bot"
  "github.com/Galagoshin/VKGoBot/bot/vk"
)

func main(){
  bot.Init()
  vk.GetHandler().RegisterResponser(Responser)
  bot.Run()
}

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
  chat.SendMessage(chats.Message{Text: "Hi, your message is " + outgoing.Text})
}
```

After the first launch, VKGoBot will create configuration files.

### Bot configuration

#### bot.gconf
```yaml
#v=1
group-tokens-list: group_tokens.list
user-tokens-list: user_tokens.list
widget-tokens-list: widget_tokens.list
group-id: 123456789
```

* `group-tokens-list` - List file of all group tokens
* `user-tokens-list` - List file of all user tokens
* `widget-tokens-list` - List file of all widget tokens
* `group-id` - Group id

### Framework configuration

#### vkgo.gconf

```yaml
#v=1
hot-reload-enabled: true
live-reload-enabled: true
payload-secret-key: secret-key
write-logs: true
debug-level: 0
```

* `live-reload-enabled` - Enable live reloads
* `hot-reload-enabled` - Enable hot reload
* `payload-secret-key` - Secret key for signing payloads in keyboards
* `write-logs` - Write logs
* `debug-level` - Set debug level

### Developer mode

Developer mode is disabled automatically when you run your application inside a docker container.

To force disable developer mode use the flag `--mode prod`

You can also enable developer mode inside a docker container with the flag `--mode dev`

### Live reloads

VKGoBot supports live reloads in developer mode. You need to correctly configure the GoLand configuration and `build.gconf`.

#### build.gconf
```yaml
#v=1
source-dir: src
run: run src/main.go
build: build -o vkgobot src/main.go
```

The `run` parameter should run your application using standart `go` utility.

#### GoLand configuration

![Example GoLand configuration](https://github.com/Galagoshin/VKGoBot/raw/master/example_conf.png)

You need to run `main()` in `github.com/Galagoshin/VKGoBot/develop/main.go`

### Hot reload

VKGoBot supports hot reload. If you change any configuration file, bot will be restarted with the new settings. You can migrate to other communities, without restart application.

## Chats

### Example responser

```go
vk.GetHandler().RegisterResponser(Responser)

func Responser(chat chats.Chat, message chats.OutgoingMessage) {
  text := message.Text //Text message from user
  user := message.User //User object, who sent message
  attachments := message.Attachments //Attachments array
  switch chat.(type) {
    case chats.UserChat:
      chat.SendMessage(chats.Message{Text: "Message from user"})
    case chats.ConversationChat:
      chat.SendMessage(chats.Message{Text: "Message from conversation"})
  }
}
```

### Attachments

#### Example creating attachment object

```go
attachment := attachments.Image{
  Id:       454535357,
  OwnerId: -123456789,
}
```

or

```go
attachment := attachments.Image{
  Id:         454535357,
  OwnerId:   -123456789,
  AccessKey: "23f17b2",
}
```

### Keyboards

#### Init keyboard

```go
kbrd := keyboards.StaticKeyboard{}
kbrd.Init()
```

#### Buttons

##### Add buttons

```go
kbrd.AddButton(keyboards.CallbackButton{
    Row: 0,
    Column: 0,
    Payload: keyboards.Payload{
        "action": "any_action",
    },
    Text: "Text",
})
```

##### Normal button

```go
kbrd.AddButton(keyboards.NormalButton{
    Row: 0,
    Column: 0,
    Color: keyboards.GreenColor,
    Payload: keyboards.Payload{
        "command": "help",
    },
    Text: "Help",
})
```

Colors in normal button

```go
Color: keyboards.GreenColor
Color: keyboards.RedColor
Color: keyboards.WhiteColor
Color: keyboards.BlueColor
```

##### Callback button

```go
kbrd.AddButton(keyboards.CallbackButton{
    Row: 0,
    Column: 0,
    Payload: keyboards.Payload{
        "action": "any_action",
    },
    Text: "Text",
})
```

Receiving callback buttons

```go
import (
	"github.com/Galagoshin/GoUtils/events"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
)

events.RegisterEvent(events.Event{Name: events2.MessageCallbackEvent, Execute: OnMessageCallbackEvent})

func OnMessageCallbackEvent(args ...any) {
    callback := args[0].(chats.Callback)
    chat := callback.Chat //Chat, where button was pressed
    payload := callback.Payload //Payload in callback button
    //...
    callback.SendAnswer(chats.CallbackAnswer{Text: "Answer in the push"})
}
```

##### Link button

```go
kbrd.AddButton(keyboards.LinkButton{
    Row: 0,
    Column: 0,
    Link: requests.URL("https://vk.com/galagoshin"),
    Text: "Open link",
})
```

#### Static keyboard

```go
kbrd := keyboards.StaticKeyboard{}
```

#### Inline keyboard

```go
kbrd := keyboards.InlineKeyboard{}
```

#### One time keyboard

```go
kbrd := keyboards.OneTimeKeyboard{}
```

#### Remove keyboard

```go
kbrd := keyboards.RemoveKeyboard{}
```

### Message

#### Example sending message with previous examples

```go
chat.SendMessage(chats.Message{
    Text:        "Text",
    Keyboard:    &kbrd,
    Attachments: &attachments,
})
```

## Plugins

VKGoBot supports plugins in `.so` files.

### Example plugin

#### Plugin code
```go
package main

var Name = "PluginName" //Required field
var Version = "1.0"     //Optional

func GetUser() int {
	return 1
}

func OnEnable(){        //Optional
	//TODO
}

func OnDisable() {      //Optional
	//TODO
}
```

#### Plugin compilation
`go build -buildmode=plugin -o plug_1.0.so`

Then move `plug_1.0.so` to the `plugins` directory.

#### Usage in VKGoBot

```go
plugin, exists := plugins.GetPlugin("PluginName", "1.0") //version can be empty to get latest version
if exists {
	symbol, err := plugin.Lookup("GetUser")
	if err != nil {
		logger.Panic(err)
	}
	GetUser := symbol.(func() int)
	return GetUser() //returns 1
}else{
	logger.Error("Plugin is not enabled.")
}
```