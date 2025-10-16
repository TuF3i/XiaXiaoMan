package helloWorld

import (
	"XiaXiaoMan/core"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tencent-connect/botgo/dto"
)

type Processor struct{}

// ProcessGroupMessage 回复群消息
func (p *Processor) ProcessGroupMessage(input string, data *dto.WSGroupATMessageData) error {
	msg := generateMessage(input, dto.Message(*data)) //生成群消息

	if err := p.sendGroupReply(context.Background(), data.GroupID, msg); err != nil {
		_ = p.sendGroupReply(context.Background(), data.GroupID, genErrMessage(dto.Message(*data), err))
	}
	return nil
}

func genErrMessage(data dto.Message, err error) *dto.MessageToCreate {
	return &dto.MessageToCreate{
		Timestamp: time.Now().UnixMilli(),
		Content:   fmt.Sprintf("处理异常:%v", err),
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
		MsgID: data.ID,
	}
}

func generateMessage(input string, data dto.Message) *dto.MessageToCreate {
	return &dto.MessageToCreate{
		Timestamp: time.Now().UnixMilli(),
		Content:   input,
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
		MsgID: data.ID,
	}
}

func (p Processor) sendGroupReply(ctx context.Context, groupID string, toCreate dto.APIMessage) error {
	log.Printf("EVENT ID:%v", toCreate.GetEventID())
	if _, err := core.Engine.PostGroupMessage(ctx, groupID, toCreate); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
