package helloWorld

import (
	"XiaXiaoMan/core"
	"context"
	"fmt"

	"github.com/tencent-connect/botgo/dto"
)

var _pluggin_Name string = "HelloWorld"

var p Processor

func (root *Body) InitThisPlugin() {
	root.PluginName = _pluggin_Name
	core.Logger.BotDEBUG(fmt.Sprintf("Init the pluggin %v", _pluggin_Name))
}

func (root *Body) ProcessEvent(data *dto.WSGroupATMessageData) error {
	content := fmt.Sprintf("Hello World!!! [%v]", data.Author.Username)
	msg := generateMessage(content, dto.Message(*data)) //生成群消息

	if err := p.sendGroupReply(context.Background(), data.GroupID, msg); err != nil {
		_ = p.sendGroupReply(context.Background(), data.GroupID, genErrMessage(dto.Message(*data), err))
	}
	return nil
}

func (root *Body) ReturnMatchCommand() string {
	return root.MatchCommand
}

func (root *Body) ReturnPluginName() string {
	return root.PluginName
}
