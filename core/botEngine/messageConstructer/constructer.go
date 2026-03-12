package messageConstructer

import "XiaXiaoMan/core/models/onebot"

func NewTextSegment(text string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "text",
		Data: map[string]interface{}{
			"text": text,
		},
	}
}

func NewImageSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "image",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func NewFaceSegment(id int) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "face",
		Data: map[string]interface{}{
			"id": id,
		},
	}
}

func NewAtSegment(userID int64) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "at",
		Data: map[string]interface{}{
			"qq": userID,
		},
	}
}

func NewRecordSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "record",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func NewVideoSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "video",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func NewFileSegment(file string) onebot.MessageSegment {
	return onebot.MessageSegment{
		Type: "file",
		Data: map[string]interface{}{
			"file": file,
		},
	}
}

func NewForwardNode(name string, uin int64, content interface{}) onebot.ForwardMessageNode {
	return onebot.ForwardMessageNode{
		Type: "node",
		Data: map[string]interface{}{
			"name":    name,
			"uin":     uin,
			"content": content,
		},
	}
}
