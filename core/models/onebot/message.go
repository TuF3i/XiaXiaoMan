package onebot

type MessageSegment struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type SendPrivateMsgRequest struct {
	UserID     int64       `json:"user_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape,omitempty"`
}

type SendPrivateForwardMsgRequest struct {
	UserID   int64                `json:"user_id"`
	Messages []ForwardMessageNode `json:"messages"`
}

type SendGroupMsgRequest struct {
	GroupID    int64       `json:"group_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape,omitempty"`
}

type SendGroupForwardMsgRequest struct {
	GroupID  int64                `json:"group_id"`
	Messages []ForwardMessageNode `json:"messages"`
}

type ForwardMessageNode struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type ForwardNode struct {
	Name    string      `json:"name"`
	Uin     int64       `json:"uin"`
	Content interface{} `json:"content"`
}

type ForwardMessage struct {
	ForwardID string `json:"forward_id"`
	NodeList  []struct {
		Sender struct {
			UserID   int64  `json:"user_id"`
			Nickname string `json:"nickname"`
		} `json:"sender"`
		Time    int64            `json:"time"`
		Message []MessageSegment `json:"message"`
	} `json:"node_list"`
}

type GetMsgRequest struct {
	MessageID int64 `json:"message_id"`
}

type DeleteMsgRequest struct {
	MessageID int64 `json:"message_id"`
}

type GetFileRequest struct {
	FileID string `json:"file_id"`
}

type GetImageRequest struct {
	File string `json:"file"`
}

type GetRecordRequest struct {
	File      string `json:"file"`
	OutFormat string `json:"out_format"`
}

type SetMsgEmojiLikeRequest struct {
	MessageID int64 `json:"message_id"`
	EmojiID   int   `json:"emoji_id"`
}

type UnsetMsgEmojiLikeRequest struct {
	MessageID int64 `json:"message_id"`
	EmojiID   int   `json:"emoji_id"`
}

type GetFriendMsgHistoryRequest struct {
	UserID     int64 `json:"user_id"`
	MessageID  int64 `json:"message_id"`
	Count      int   `json:"count,omitempty"`
	MessageSeq int64 `json:"message_seq,omitempty"`
}

type GetGroupMsgHistoryRequest struct {
	GroupID    int64 `json:"group_id"`
	MessageID  int64 `json:"message_id"`
	Count      int   `json:"count,omitempty"`
	MessageSeq int64 `json:"message_seq,omitempty"`
}

type GetForwardMsgRequest struct {
	ID string `json:"id"`
}

type MarkMsgAsReadRequest struct {
	MessageID int64 `json:"message_id"`
}

type VoiceMsgToTextRequest struct {
	MessageID int64  `json:"message_id"`
	Voice     string `json:"voice,omitempty"`
}

type SendGroupAIRecordRequest struct {
	GroupID   int64  `json:"group_id"`
	Character string `json:"character,omitempty"`
	Text      string `json:"text"`
}

type GetAICharactersRequest struct {
	Type string `json:"type,omitempty"`
}

type ForwardFriendSingleMsgRequest struct {
	UserID    int64 `json:"user_id"`
	MessageID int64 `json:"message_id"`
}

type ForwardGroupSingleMsgRequest struct {
	GroupID   int64 `json:"group_id"`
	MessageID int64 `json:"message_id"`
}

type MessageInfo struct {
	Time        int64            `json:"time"`
	MessageType string           `json:"message_type"`
	MessageID   int64            `json:"message_id"`
	RealID      int64            `json:"real_id"`
	Sender      MessageSender    `json:"sender"`
	Message     []MessageSegment `json:"message"`
	RawMessage  string           `json:"raw_message"`
}

type MessageSender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Card     string `json:"card,omitempty"`
	Area     string `json:"area,omitempty"`
	Level    string `json:"level,omitempty"`
	Role     string `json:"role,omitempty"`
	Title    string `json:"title,omitempty"`
}

type FileInfo struct {
	File     string `json:"file"`
	URL      string `json:"url"`
	FileSize string `json:"file_size"`
}

type ImageInfo struct {
	File     string `json:"file"`
	URL      string `json:"url"`
	Size     int    `json:"size"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filename string `json:"filename"`
}

type RecordInfo struct {
	File string `json:"file"`
	URL  string `json:"url"`
}

type MsgHistory struct {
	Messages []MessageInfo `json:"messages"`
}

type VoiceToText struct {
	Text string `json:"text"`
}

type AICharacter struct {
	CharacterID     string `json:"character_id"`
	CharacterName   string `json:"character_name"`
	CharacterAvatar string `json:"character_avatar"`
}

type AICharacters struct {
	Characters []AICharacter `json:"characters"`
}

type MsgSendResponse struct {
	MessageID int64 `json:"message_id"`
}
