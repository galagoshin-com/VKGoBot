package tasks

import (
	"errors"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/crypto"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/GoUtils/scheduler"
	"github.com/Galagoshin/VKGoBot/bot/framework"
	"os"
	"path/filepath"
	"time"
)

var HotReloadTask = &scheduler.RepeatingTask{
	Duration:   time.Second,
	OnComplete: HotReloadExecutor,
}

var lastFileHash = make(map[string]string)

func HotReloadExecutor(args ...any) {
	task := args[0].(*scheduler.RepeatingTask)
	all_configs, err := filepath.Glob("./*.gconf")
	if err != nil {
		logger.Error(err)
		task.Destroy()
	}
	for _, config := range all_configs {
		file := files.File{Path: config}
		err := file.Open(os.O_RDWR)
		if err != nil {
			logger.Error(err)
			task.Destroy()
		}
		hash := crypto.Sha1(file.Read())
		lastHash, exists := lastFileHash[config]
		if !exists || lastHash != hash {
			lastFileHash[config] = hash
			prevLive, exists := framework.Config.Get("live-reload-enabled")
			if !exists {
				logger.Panic(errors.New("\"live-reload-enabled\" was removed from the framework config."))
			}
			prevHot, exists := framework.Config.Get("hot-reload-enabled")
			if !exists {
				logger.Panic(errors.New("\"hot-reload-enabled\" was removed from the framework config."))
			}
			framework.HotReload()
			hotreload, hotreloadError := framework.Config.Get("hot-reload-enabled")
			if !hotreloadError {
				logger.Panic(errors.New("\"hot-reload-enabled\" is not defined in the framework config."))
			}
			if hotreload.(string) == "false" && prevHot.(string) != hotreload.(string) {
				task.Destroy()
			}
			livereload, livereloadError := framework.Config.Get("live-reload-enabled")
			if !livereloadError {
				logger.Panic(errors.New("\"live-reload-enabled\" is not defined in the framework config."))
			}
			if livereload.(string) == "false" && prevLive.(string) != livereload.(string) {
				RestartTask.Destroy()
			} else if livereload.(string) == "true" && prevLive.(string) != livereload.(string) {
				InitRestartTask()
				RestartTask.Run(RestartTask)
			}
			return
		}
	}
}

func InitHotReloadTask() {
	all_configs, err := filepath.Glob("./*.gconf")
	if err != nil {
		logger.Panic(err)
	}
	for _, config := range all_configs {
		file := files.File{Path: config}
		err := file.Open(os.O_RDWR)
		if err != nil {
			logger.Panic(err)
		}
		lastFileHash[config] = crypto.Sha1(file.Read())
	}
}
