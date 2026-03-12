package botEngine

import (
	"XiaXiaoMan/core/models/onebot"
)

// SetMsgEmojiLike 给消息添加表情
func (c *BotEngine) SetMsgEmojiLike(messageID int64, emojiID int) (err error) {
	req := onebot.SetMsgEmojiLikeRequest{
		MessageID: messageID,
		EmojiID:   emojiID,
	}
	_, err = c.SendRequest("set_msg_emoji_like", req)
	return err
}
