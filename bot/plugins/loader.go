package plugins

import (
	"errors"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/events"
	"github.com/Galagoshin/GoUtils/files"
	events2 "github.com/Galagoshin/VKGoBot/bot/events"
	"path/filepath"
	"plugin"
)

func EnableAllPlugins() {
	pluginsDir := files.Directory{Path: "plugins"}
	err := pluginsDir.CreateAll()
	if err != nil {
		logger.Panic(err)
	}
	all_plugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		logger.Panic(err)
	}

	for _, filename := range all_plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			logger.Panic(err)
		}

		symbolPluginName, err := p.Lookup("Name")
		pluginName, ok := symbolPluginName.(*string)
		if err != nil || !ok {
			logger.Panic(errors.New("Plugin has no \"Name\" const"))
		}

		symbolPluginVersion, err := p.Lookup("Version")
		pluginVersion, ok := symbolPluginVersion.(*string)
		if err != nil || !ok {
			logger.Warning("Plugin has no \"Version\" const")
		}

		symbolOnEnable, err := p.Lookup("OnEnable")
		onEnable, ok := symbolOnEnable.(func())
		if err != nil || !ok {
			logger.Warning("Plugin has no \"OnEnable()\" function")
		}

		symbolOnDisable, err := p.Lookup("OnDisable")
		onDisable, ok := symbolOnDisable.(func())
		if err != nil || !ok {
			logger.Warning("Plugin has no \"OnDisable()\" function")
		}

		plug := Plugin{
			Name:      *pluginName,
			Version:   *pluginVersion,
			OnEnable:  onEnable,
			OnDisable: onDisable,
			plugin:    p,
		}
		_, exists := pluginStorage[*pluginName]
		if exists {
			pluginStorage[*pluginName][*pluginVersion] = &plug
		} else {
			pluginStorage[*pluginName] = make(map[string]*Plugin)
			pluginStorage[*pluginName][*pluginVersion] = &plug
		}
	}
	pluginsCounter := 0
	for _, vers := range pluginStorage {
		for _, plug := range vers {
			if plug.OnEnable != nil {
				plug.OnEnable()
			}
			events.CallAllEvents(events2.EnablePluginEvent, plug)
			version := ""
			if plug.Version != "" {
				version = "v" + plug.Version
			}
			logger.Print(fmt.Sprintf("Plugin %s %s has been enabled.", plug.Name, version))
			pluginsCounter++
		}
	}
	logger.Print(fmt.Sprintf("%d plugins has been enabled.", pluginsCounter))
}

func DisableAllPlugins() {
	pluginsCounter := 0
	for _, vers := range pluginStorage {
		for _, plug := range vers {
			if plug.OnDisable != nil {
				plug.OnDisable()
			}
			events.CallAllEvents(events2.DisablePluginEvent, plug)
			version := ""
			if plug.Version != "" {
				version = "v" + plug.Version
			}
			pluginsCounter++
			logger.Print(fmt.Sprintf("Plugin %s %s has been disabled.", plug.Name, version))
		}
	}
	logger.Print(fmt.Sprintf("%d plugins has been disabled.", pluginsCounter))
}
