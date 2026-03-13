package botEngine

import (
	"XiaXiaoMan/core/config"
	"XiaXiaoMan/core/models/onebot"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	PrivateMessageEvent  = "event.message.private"       // 私聊消息事件
	GroupMessageEvent    = "event.message.group"         // 群聊消息事件
	GroupFileUploadEvent = "event.notice.group.upload "  // 群文件上传事件
	GroupAdminEvent      = "event.notice.group.admin"    // 群管理员变动事件
	GroupDecreaseEvent   = "event.notice.group.decrease" // 群成员减少事件
	GroupIncreaseEvent   = "event.notice.group.increase" // 群成员增加事件
	GroupBanEvent        = "event.notice.group.ban"      // 群成员禁言事件
	GroupRecallEvent     = "event.notice.group.recall"   // 群消息撤回事件
	GroupCardEvent       = "event.notice.group.card"     // 群名片变更事件
	FriendAddEvent       = "event.notice.friend.add"     // 好友添加事件
	FriendRecallEvent    = "event.notice.friend.recall"  // 好友消息撤回事件
	NotifyEvent          = "event.notice.notify"         // 特殊通知事件（如戳一戳、红包王等）
	OfflineFileEvent     = "event.notice.file.offline"   // 离线文件事件
	EssenceEvent         = "event.notice.essence"        // 群精华消息事件
	FriendEvent          = "event.request.friend"        // 好友请求事件
	GroupEvent           = "event.request.group"         // 群请求事件
	LifecycleMetaEvent   = "event.meta.lifecycle"        // 生命周期事件
	Heartbeat            = "event.meta.heartbeat"        // 心跳事件
)

type BotEngine struct {
	conf *config.Config
	conn *websocket.Conn
	uuid uuid.UUID
	mu   sync.Mutex
	wg   sync.WaitGroup

	pendingCalls map[string]chan *onebot.APIResponse
	eventChan    chan interface{}
	closeChan    chan struct{}

	handlerFunc map[string]HandleFunc
}

type RequestContext struct {
	c              *BotEngine
	requestContext interface{}
	customContext  map[string]interface{}
}

func genRequestContext(c *BotEngine, requestContext interface{}) *RequestContext {
	return &RequestContext{
		c:              c,
		requestContext: requestContext,
		customContext:  make(map[string]interface{}),
	}
}
