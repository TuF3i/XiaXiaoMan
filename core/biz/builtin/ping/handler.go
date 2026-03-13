package ping

import (
	"XiaXiaoMan/core/botEngine"
	"XiaXiaoMan/core/models/onebot"
	"context"
	"fmt"
)

func (r *Command) Ping(ctx context.Context, c *botEngine.RequestContext) {
	req := c.GetRequestContext()
	if data, ok := req.(onebot.GroupMessageEvent); ok {
		msg := c.NewMessageSegmentSet().Add(c.NewAtSegment(data.UserID)).Add(c.NewTextSegment("pong")).Build()
		msgID, err := c.SendGroupMessage(data.GroupID, msg, false)
		if err != nil {
			fmt.Printf("[Error] Message %v error: %v", msgID, err.Error())
			return
		}
	}

	if data, ok := req.(onebot.PrivateMessageEvent); ok {
		msg := c.NewMessageSegmentSet().Add(c.NewTextSegment("pong")).Build()
		msgID, err := c.SendPrivateMessage(data.UserID, msg, false)
		if err != nil {
			fmt.Printf("[Error] Message %v error: %v", msgID, err.Error())
			return
		}
	}
}
