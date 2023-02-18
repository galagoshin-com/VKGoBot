package commands

import "github.com/Galagoshin/VKGoBot/bot/framework"

func Stop(string, []string) {
	framework.Shutdown(false)
}
