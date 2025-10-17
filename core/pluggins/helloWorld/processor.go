package helloWorld

import (
	"context"
	"fmt"

	"github.com/tencent-connect/botgo/dto"
)

var p Processor

func (root *HelloWorld) ProcessEvent(data *dto.WSGroupATMessageData) error {
	content := fmt.Sprintf("Hello World!!! [%v]", data.Author.Username)
	msg := generateMessage(content, dto.Message(*data)) //生成群消息

	if err := p.sendGroupReply(context.Background(), data.GroupID, msg); err != nil {
		_ = p.sendGroupReply(context.Background(), data.GroupID, genErrMessage(dto.Message(*data), err))
	}
	return nil
}

func (root *HelloWorld) ReturnMetchCommand() string {
	return root.MatchCommand
}
