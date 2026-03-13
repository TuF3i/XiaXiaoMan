package botEngine

import (
	"XiaXiaoMan/core/models/onebot"
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

// SendPrivateMessage 发送私聊消息
func (r *RequestContext) SendPrivateMessage(userID int64, message interface{}, autoEscape bool) (messageID int64, err error) {
	// 构造请求
	req := onebot.SendPrivateMsgRequest{
		UserID:     userID,
		Message:    message,
		AutoEscape: autoEscape,
	}
	// 向LLBot发起调用
	resp, err := r.c.sendRequest("send_private_msg", req)
	if err != nil {
		return 0, err
	}
	// 解析响应
	dataBytes, _ := sonic.Marshal(resp.Data)
	var result onebot.MsgSendResponse
	if err := sonic.Unmarshal(dataBytes, &result); err != nil {
		return 0, fmt.Errorf("反序列化数据错误: %v", err.Error())
	}

	return result.MessageID, nil
}

// SendPrivateTextMessage 发送私聊纯文本消息
func (r *RequestContext) SendPrivateTextMessage(userID int64, message string) (messageID int64, err error) {
	msg := []onebot.MessageSegment{
		{
			Type: "text",
			Data: map[string]interface{}{
				"text": message,
			},
		},
	}
	return r.SendPrivateMessage(userID, msg, false)
}

// SendPrivateForwardMessage 发送私聊合并转发消息
func (r *RequestContext) SendPrivateForwardMessage(userID int64, messages []onebot.ForwardMessageNode) (forwardID string, err error) {
	// 构造请求
	req := onebot.SendPrivateForwardMsgRequest{
		UserID:   userID,
		Messages: messages,
	}
	// 发起调用
	resp, err := r.c.sendRequest("send_private_forward_msg", req)
	if err != nil {
		return "", err
	}
	dataBytes, _ := json.Marshal(resp.Data)
	var result map[string]string
	_ = sonic.Unmarshal(dataBytes, &result)
	return result["forward_id"], nil
}

// GetFriendMsgHistory 获取私聊的聊天记录
func (r *RequestContext) GetFriendMsgHistory(userID int64, messageID int64, count int, messageSeq int64) (histories *onebot.MsgHistory, err error) {
	req := onebot.GetFriendMsgHistoryRequest{
		UserID:     userID,
		MessageID:  messageID,
		Count:      count,
		MessageSeq: messageSeq,
	}
	resp, err := r.c.sendRequest("get_friend_msg_history", req)
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
