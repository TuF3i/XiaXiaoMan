package plugins

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/plugins/helloWorld"
	"fmt"
	"strings"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

var plugins []plugin = []plugin{
	&helloWorld.Body{MatchCommand: "/test"},
}

type plugin interface {
	ProcessEvent(data *dto.WSGroupATMessageData) error
	ReturnMatchCommand() string
	ReturnPluginName() string
	InitThisPlugin()
}

func InitPlugins() {
	registedPlugins := make([]string, 0)
	for _, p := range plugins {
		p.InitThisPlugin()
		registedPlugins = append(registedPlugins, p.ReturnPluginName())
		core.Logger.BotDEBUG(fmt.Sprintf("Init the pluggin %v", registedPlugins))
	}
}

func ProcessGroupMessage(data *dto.WSGroupATMessageData) error {
	raw := strings.ToLower(message.ETLInput(data.Content))
	commandSlice := strings.SplitN(raw, " ", 2)
	cmd := commandSlice[0]

	for _, p := range plugins {
		if cmd == p.ReturnMatchCommand() {
			p.ProcessEvent(data)
		}
	}

	return nil
}
