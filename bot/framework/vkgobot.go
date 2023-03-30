package framework

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/configs"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/time"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
	"github.com/Galagoshin/VKGoBot/bot/plugins"
	"github.com/Galagoshin/VKGoBot/bot/vk"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/sign"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type RuntimeMode uint8

const (
	ProductionMode = RuntimeMode(iota)
	DevelopMode    = RuntimeMode(iota)
)

var Mode RuntimeMode

var Config = &configs.Config{Name: "vkgo"}
var BuildConfig = &configs.Config{Name: "build"}

func Build(err_out error) {
	build_str, exists := BuildConfig.Get("build")
	if !exists {
		logger.Panic(errors.New("\"run\" is not defined in the build config."))
	}
	cmd := exec.Command("go", strings.Split(build_str.(string), " ")...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		err_out = err
		return
	}
	cmd.Stdin = os.Stdin
	err = cmd.Start()
	if err != nil {
		err_out = err
		return
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		logger.Print(m)
	}
	err = cmd.Wait()
	if err != nil {
		err_out = err
		return
	}
}

func IsUnderDocker() bool {
	file := files.File{Path: "/.dockerenv"}
	return file.Exists()
}

func HotReload() {
	logger.Print(fmt.Sprintf("HotReload finished (%f s.)", time.MeasureExecution(func() {

		plugins.DisableAllPlugins()
		Config = &configs.Config{Name: "vkgo"}
		BuildConfig.Init(map[string]any{
			"run":        "run src/main.go",
			"build":      "build -o vkgobot src/main.go",
			"source-dir": "src",
		}, 2)
		runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ%!@#$^&*()_+-0987654321")
		secret := make([]rune, 128)
		for i := range secret {
			secret[i] = runes[rand.Intn(len(runes))]
		}
		Config.Init(map[string]any{
			"hot-reload-enabled":  "true",
			"live-reload-enabled": "true",
			"payload-secret-key":  string(secret),
			"write-logs":          "true",
			"debug-level":         0,
		}, 1)
		vk.Config = &configs.Config{Name: "bot"}

		vk.Stop()
		events.CallAllEvents(events2.HotReloadEvent)

		secretKey, secretError := Config.Get("payload-secret-key")
		if !secretError {
			logger.Panic(errors.New("\"payload-secret-key\" is not defined in the framework config."))
		}

		debugLevel, debugError := Config.Get("debug-level")
		if !debugError {
			logger.Panic(errors.New("\"debug-level\" is not defined in the framework config."))
		}

		writeLogs, writeLogsError := Config.Get("write-logs")
		if !writeLogsError {
			logger.Panic(errors.New("\"write-logs\" is not defined in the framework config."))
		}

		logger.SetDebugLevel(debugLevel.(int))
		logger.SetLogs(writeLogs.(string) == "true")

		sign.SECRET = secretKey.(string)

		plugins.EnableAllPlugins()

		vk.Init()
		go vk.Run()
	})))
}

func Shutdown(restart bool) {
	if restart {
		logger.Print("Framework is reloading...")
	} else {
		logger.Print("Framework is shuting down...")
	}
	events.CallAllEvents(events2.StopApplicationEvent)
	plugins.DisableAllPlugins()
	err := vk.Config.Save()
	if err != nil {
		logger.Error(err)
	}
	if restart {
		os.Exit(0)
	}
	os.Exit(130)
}
