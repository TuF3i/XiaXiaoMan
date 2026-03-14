package botEngine

import (
	"XiaXiaoMan/core/models/onebot"
	"context"
)

type HandleFunc func(ctx context.Context, c *RequestContext)

func (r *RequestContext) GetRequestContext() interface{} {
	return r.requestContext
}

func (r *RequestContext) GetGroupMessageEvent() onebot.GroupMessageEvent {
	return r.requestContext.(onebot.GroupMessageEvent)
}

func (r *RequestContext) GetFriendAddNoticeEvent() onebot.FriendAddNoticeEvent {
	return r.requestContext.(onebot.FriendAddNoticeEvent)
}

func (r *RequestContext) GetGroupIncreaseNoticeEvent() onebot.GroupIncreaseNoticeEvent {
	return r.requestContext.(onebot.GroupIncreaseNoticeEvent)
}

func (r *RequestContext) GetGroupDecreaseNoticeEvent() onebot.GroupDecreaseNoticeEvent {
	return r.requestContext.(onebot.GroupDecreaseNoticeEvent)
}

func (c *BotEngine) RegisterEventHandlerFunc(eventType string, handlerFunc HandleFunc) {
	c.handlerFunc[eventType] = handlerFunc
}
