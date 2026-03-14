package ping

import (
	"XiaXiaoMan/core/botEngine"
	"context"
	"fmt"
)

func (r *Command) Ping(ctx context.Context, c *botEngine.RequestContext) {
	req := c.GetGroupMessageEvent()
	msg := c.NewMessageSegmentSet().Add(c.NewAtSegment(req.UserID)).Add(c.NewTextSegment("pong")).Build()
	msgID, err := c.SendGroupMessage(req.GroupID, msg, false)
	if err != nil {
		fmt.Printf("[Error] Message %v error: %v", msgID, err.Error())
		return
	}
}
