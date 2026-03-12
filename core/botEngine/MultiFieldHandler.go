package botEngine

import (
	"XiaXiaoMan/core/models/onebot"

	"github.com/bytedance/sonic"
)

// GetFile 获取文件内容
func (c *BotEngine) GetFile(fileID string) (fileData *onebot.FileInfo, err error) {
	req := onebot.GetFileRequest{
		FileID: fileID,
	}
	resp, err := c.SendRequest("get_file", req)
	if err != nil {
		return nil, err
	}
	dataBytes, _ := sonic.Marshal(resp.Data)
	var file onebot.FileInfo
	if err := sonic.Unmarshal(dataBytes, &file); err != nil {
		return nil, err
	}
	return &file, nil
}

// GetImage 获取图片内容
func (c *BotEngine) GetImage(file string) (imageData *onebot.ImageInfo, err error) {
	req := onebot.GetImageRequest{
		File: file,
	}
	resp, err := c.SendRequest("get_image", req)
	if err != nil {
		return nil, err
	}
	dataBytes, _ := sonic.Marshal(resp.Data)
	var image onebot.ImageInfo
	if err := sonic.Unmarshal(dataBytes, &image); err != nil {
		return nil, err
	}
	return &image, nil
}

// VoiceMsgToText 语言转文字
func (c *BotEngine) VoiceMsgToText(messageID int64, voice string) (text *onebot.VoiceToText, err error) {
	req := onebot.VoiceMsgToTextRequest{
		MessageID: messageID,
		Voice:     voice,
	}
	resp, err := c.SendRequest("voice_msg_to_text", req)
	if err != nil {
		return nil, err
	}
	dataBytes, _ := sonic.Marshal(resp.Data)
	var result onebot.VoiceToText
	if err := sonic.Unmarshal(dataBytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
