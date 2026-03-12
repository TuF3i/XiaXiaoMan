package eventHandler

import (
	"XiaXiaoMan/core/botEngine"
	"XiaXiaoMan/core/models/onebot"
	"encoding/json"

	"github.com/bytedance/sonic"
)

// GetMsg 获取消息内容
func GetMsg(c *botEngine.BotEngine, messageID int64) (message *onebot.MessageInfo, err error) {
	req := onebot.GetMsgRequest{
		MessageID: messageID,
	}
	resp, err := c.SendRequest("get_msg", req)
	if err != nil {
		return nil, err
	}
	dataBytes, _ := json.Marshal(resp.Data)
	var msg onebot.MessageInfo
	if err := sonic.Unmarshal(dataBytes, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// DeleteMsg 撤回消息
func DeleteMsg(c *botEngine.BotEngine, messageID int64) (err error) {
	req := onebot.DeleteMsgRequest{
		MessageID: messageID,
	}
	_, err = c.SendRequest("delete_msg", req)
	return err
}
