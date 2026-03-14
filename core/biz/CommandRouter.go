package biz

import (
	"XiaXiaoMan/core/botEngine"
	"context"
	"strings"
)

func detectPrefix(msg string) bool {
	if []byte(msg)[0] == '/' {
		return true
	}
	return false
}

func extractCommandCaller(msg string) string {
	if len(msg) == 0 {
		return ""
	}

	cmdWithPrefixBytes := []byte(strings.Split(msg, " ")[0])

	if len(cmdWithPrefixBytes) == 0 {
		return ""
	}

	rawCmd := cmdWithPrefixBytes[1:len(cmdWithPrefixBytes)]

	return string(rawCmd)
}

func (r *Biz) BizRouter(ctx context.Context, c *botEngine.RequestContext) {
	msg := c.GetGroupMessageEvent()
	if detectPrefix(msg.RawMessage) {
		cmdCaller := extractCommandCaller(msg.RawMessage)
		cmd, ok := r.GetBuiltinCommand(cmdCaller)
		if ok {
			cmd.GetHandlerFunc()(ctx, c)
		}

		cmd, ok = r.GetMagicCommand(cmdCaller)
		if ok {
			cmd.GetHandlerFunc()(ctx, c)
		}

		cmd, ok = r.GetLuaPlugin(cmdCaller)
		if ok {
			cmd.GetHandlerFunc()(ctx, c)
		}

		_, _ = c.SendGroupMessage(
			msg.GroupID,
			c.NewMessageSegmentSet().Add(c.NewTextSegment("找不到命令喵(┬┬﹏┬┬)")).Build(),
			false,
		)
	}

	// TODO AI回复
}
