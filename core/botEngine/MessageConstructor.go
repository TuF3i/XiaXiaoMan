package botEngine

import "XiaXiaoMan/core/models/onebot"

type MessageSegmentSet struct {
	msg []onebot.MessageSegment
}

func (r *RequestContext) NewTextSegment(text string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "text",
		Data: map[string]interface{}{
			"text": text,
		},
	}
}

func (r *RequestContext) NewImageSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "image",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func (r *RequestContext) NewFaceSegment(id int) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "face",
		Data: map[string]interface{}{
			"id": id,
		},
	}
}

func (r *RequestContext) NewAtSegment(userID int64) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "at",
		Data: map[string]interface{}{
			"qq": userID,
		},
	}
}

func (r *RequestContext) NewRecordSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "record",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func (r *RequestContext) NewVideoSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "video",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func (r *RequestContext) NewFileSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "file",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func (r *RequestContext) NewForwardNode(name string, uin int64, content interface{}) onebot.ForwardMessageNode {
	return onebot.ForwardMessageNode{
		Type: "node",
		Data: map[string]interface{}{
			"name":    name,
			"uin":     uin,
			"content": content,
		},
	}
}

func (r *RequestContext) NewMessageSegmentSet() *MessageSegmentSet {
	return &MessageSegmentSet{msg: make([]onebot.MessageSegment, 0)}
}

func (m *MessageSegmentSet) Add(item onebot.MessageSegment) *MessageSegmentSet {
	m.msg = append(m.msg, item)
	return m
}

func (m *MessageSegmentSet) Build() []onebot.MessageSegment {
	return m.msg
}
