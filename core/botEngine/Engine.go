package botEngine

import (
	"XiaXiaoMan/core/models/onebot"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
)

func (c *BotEngine) connectWS() error {
	// 生成URL
	wsURL, err := url.Parse(c.conf.LLBotConfig.Url)
	if err != nil {
		return fmt.Errorf("parse url error: %v", err.Error())
	}
	// 设置头部
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", c.conf.LLBotConfig.Token))
	// 创建连接
	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), headers)
	if err != nil {
		return fmt.Errorf("dialer connection error: %v", err.Error())
	}

	c.conn = conn
	return nil
}

func (c *BotEngine) readLoop() {
	for {
		select {
		case <-c.closeChan:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				select {
				case <-c.closeChan:
				default:
					fmt.Printf("Read Message Error: %v\n", err)
				}
				return
			}
			c.handleMessage(message)
		}
	}
}

func (c *BotEngine) handleMessage(data []byte) {
	var base map[string]interface{}
	if err := sonic.Unmarshal(data, &base); err != nil {
		fmt.Printf("解析消息失败: %v\n", err)
		return
	}
	// 若是操作响应
	if echo, ok := base["echo"].(string); ok && echo != "" {
		c.mu.Lock()
		ch, exists := c.pendingCalls[echo]
		if exists {
			delete(c.pendingCalls, echo)
		}
		c.mu.Unlock()
		if exists {
			var resp onebot.APIResponse
			if err := json.Unmarshal(data, &resp); err == nil {
				ch <- &resp
			}
			close(ch)
		}
		return
	}
	// 若是QQ事件
	if postType, ok := base["post_type"].(string); ok {
		c.dispatchEvent(postType, data)
	}
}

func (c *BotEngine) dispatchEvent(postType string, data []byte) {
	var event interface{}
	switch postType {
	case "message":
		var msgEvent onebot.MessageEvent
		if err := json.Unmarshal(data, &msgEvent); err == nil {
			if msgEvent.MessageType == "private" {
				var privEvent onebot.PrivateMessageEvent
				if json.Unmarshal(data, &privEvent) == nil {
					event = privEvent
				}
			} else if msgEvent.MessageType == "group" {
				var groupEvent onebot.GroupMessageEvent
				if json.Unmarshal(data, &groupEvent) == nil {
					event = groupEvent
				}
			}
		}
	case "notice":
		var noticeEvent onebot.NoticeEvent
		if err := json.Unmarshal(data, &noticeEvent); err == nil {
			switch noticeEvent.NoticeType {
			case "group_upload":
				var e onebot.GroupUploadNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_admin":
				var e onebot.GroupAdminNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_decrease":
				var e onebot.GroupDecreaseNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_increase":
				var e onebot.GroupIncreaseNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_ban":
				var e onebot.GroupBanNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "friend_add":
				var e onebot.FriendAddNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_recall":
				var e onebot.GroupRecallNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "friend_recall":
				var e onebot.FriendRecallNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "notify":
				var e onebot.NotifyNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "group_card":
				var e onebot.GroupCardNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "offline_file":
				var e onebot.OfflineFileNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			case "essence":
				var e onebot.EssenceNoticeEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			}
		}
	case "request":
		var requestEvent onebot.RequestEvent
		if err := json.Unmarshal(data, &requestEvent); err == nil {
			if requestEvent.RequestType == "friend" {
				var e onebot.FriendRequestEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			} else if requestEvent.RequestType == "group" {
				var e onebot.GroupRequestEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			}
		}
	case "meta_event":
		var metaEvent onebot.MetaEvent
		if err := json.Unmarshal(data, &metaEvent); err == nil {
			if metaEvent.MetaEventType == "lifecycle" {
				var e onebot.LifecycleMetaEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			} else if metaEvent.MetaEventType == "heartbeat" {
				var e onebot.HeartbeatMetaEvent
				if json.Unmarshal(data, &e) == nil {
					event = e
				}
			}
		}
	}
	if event != nil {
		select {
		case c.eventChan <- event:
		default:
		}
	}
}

func (c *BotEngine) SendRequest(action string, params interface{}) (*onebot.APIResponse, error) {
	c.mu.Lock()
	if c.conn == nil {
		c.mu.Unlock()
		return nil, fmt.Errorf("WebSocket未连接")
	}
	echo := fmt.Sprintf("echo_%s", c.uuid.String()[:8])
	respChan := make(chan *onebot.APIResponse, 1)
	c.pendingCalls[echo] = respChan
	req := onebot.APIRequest{
		Action: action,
		Params: params,
		Echo:   echo,
	}
	data, err := json.Marshal(req)
	if err != nil {
		c.mu.Unlock()
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}
	c.mu.Unlock()
	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	select {
	case resp := <-respChan:
		if resp.Retcode != 0 {
			return resp, fmt.Errorf("API调用失败: retcode=%d, message=%s", resp.Retcode, resp.Message)
		}
		return resp, nil
	case <-time.After(30 * time.Second):
		c.mu.Lock()
		delete(c.pendingCalls, echo)
		c.mu.Unlock()
		return nil, fmt.Errorf("请求超时")
	case <-c.closeChan:
		return nil, fmt.Errorf("连接已关闭")
	}
}

func (c *BotEngine) EventChannel() <-chan interface{} {
	return c.eventChan
}
