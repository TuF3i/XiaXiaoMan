package botEngine

import "context"

const (
	PrivateMessageEvent  = "event.message.private" // 私聊消息事件
	GroupMessageEvent    = "event.message.group"   // 群聊消息事件
	GroupFileUploadEvent = "event.group.upload "   // 群文件上传事件
	GroupAdminEvent      = "event.group.admin"     // 群管理员变动事件
	
)

type HandleFunc func(ctx context.Context, c *BotEngine)

func (c *BotEngine) RegisterEventHandlerFunc(eventType string, handlerFunc HandleFunc) {
	c.handlerFunc[eventType] = handlerFunc
}
