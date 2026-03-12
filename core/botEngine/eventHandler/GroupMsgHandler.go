package eventHandler

import (
	"XiaXiaoMan/core/botEngine"
	"XiaXiaoMan/core/models/onebot"
	"encoding/json"

	"github.com/bytedance/sonic"
)

// SendGroupMessage 发送群聊消息
func SendGroupMessage(c *botEngine.BotEngine, groupID int64, message interface{}, autoEscape bool) (messageID int64, err error) {
	req := onebot.SendGroupMsgRequest{
		GroupID:    groupID,
		Message:    message,
		AutoEscape: autoEscape,
	}
	resp, err := c.SendRequest("send_group_msg", req)
	if err != nil {
		return 0, err
	}
	dataBytes, _ := sonic.Marshal(resp.Data)
	var result onebot.MsgSendResponse
	if err := sonic.Unmarshal(dataBytes, &result); err != nil {
		return 0, err
	}
	return result.MessageID, nil
}

// SendGroupTextMessage 发送群文本消息
func SendGroupTextMessage(c *botEngine.BotEngine, groupID int64, message string) (messageID int64, err error) {
	msg := []onebot.MessageSegment{
		{
			Type: "text",
			Data: map[string]interface{}{
				"text": message,
			},
		},
	}
	return SendGroupMessage(c, groupID, msg, false)
}

// SendGroupForwardMsg 发送群聊合并转发消息
func SendGroupForwardMsg(c *botEngine.BotEngine, groupID int64, messages []onebot.ForwardMessageNode) (forwardID string, err error) {
	req := onebot.SendGroupForwardMsgRequest{
		GroupID:  groupID,
		Messages: messages,
	}
	resp, err := c.SendRequest("send_group_forward_msg", req)
	if err != nil {
		return "", err
	}
	dataBytes, _ := json.Marshal(resp.Data)
	var result map[string]string
	_ = sonic.Unmarshal(dataBytes, &result)
	return result["forward_id"], nil
}

// GetGroupMsgHistory 获取群聊历史记录
func GetGroupMsgHistory(c *botEngine.BotEngine, groupID int64, messageID int64, count int, messageSeq int64) (histories *onebot.MsgHistory, err error) {
	req := onebot.GetGroupMsgHistoryRequest{
		GroupID:    groupID,
		MessageID:  messageID,
		Count:      count,
		MessageSeq: messageSeq,
	}
	resp, err := c.SendRequest("get_group_msg_history", req)
	if err != nil {
		return nil, err
	}
	dataBytes, _ := sonic.Marshal(resp.Data)
	var history onebot.MsgHistory
	if err := sonic.Unmarshal(dataBytes, &history); err != nil {
		return nil, err
	}
	return &history, nil
}
