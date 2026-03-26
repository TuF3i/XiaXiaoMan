package llm

import (
	"XiaXiaoMan/core/models/onebot"
	"fmt"
	"strings"
)

func (r *Replier) MsgExtractor(msg []onebot.MessageSegment) string {
	var msgString strings.Builder
	if len(msg) == 0 {
		return "这条消息为空"
	}

	for _, segment := range msg {
		switch segment.Type {
		case "text":
			item, ok := segment.Data["text"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(item)
		case "face":
			item, ok := segment.Data["id"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[QQ表情]id:%s[/QQ表情]", item))
		case "image":
			item, ok := segment.Data["file"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[QQ图片]url:%s[/QQ图片]", item))
		case "share":
			url, ok := segment.Data["url"].(string)
			if !ok {
				continue
			}
			title, ok := segment.Data["title"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[QQ分享链接]title:%s,url:%s[/QQ分享链接]", title, url))
		case "location":
			lat, ok := segment.Data["lat"].(string)
			if !ok {
				continue
			}
			lon, ok := segment.Data["lon"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[QQ位置]lat:%s,lon:%s[/QQ位置]", lat, lon))
		case "reply":
			id, ok := segment.Data["id"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[回复时引用的消息ID]id:%s[/回复时引用的消息ID]", id))
		case "forward":
			id, ok := segment.Data["id"].(string)
			if !ok {
				continue
			}
			msgString.WriteString(fmt.Sprintf("[合并转发消息的ID]id:%s[/合并转发消息的ID]", id))
		}
	}

	if msgString.String() == "" {
		return "这条消息为空"
	}

	return msgString.String()
}
