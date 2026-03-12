package onebot

type PostType string

const (
	PostTypeMessage   PostType = "message"
	PostTypeNotice    PostType = "notice"
	PostTypeRequest   PostType = "request"
	PostTypeMetaEvent PostType = "meta_event"
)

type MessageType string

const (
	MessageTypePrivate MessageType = "private"
	MessageTypeGroup   MessageType = "group"
)

type NoticeType string

const (
	NoticeTypeGroupUpload   NoticeType = "group_upload"
	NoticeTypeGroupAdmin    NoticeType = "group_admin"
	NoticeTypeGroupDecrease NoticeType = "group_decrease"
	NoticeTypeGroupIncrease NoticeType = "group_increase"
	NoticeTypeGroupBan      NoticeType = "group_ban"
	NoticeTypeFriendAdd     NoticeType = "friend_add"
	NoticeTypeGroupRecall   NoticeType = "group_recall"
	NoticeTypeFriendRecall  NoticeType = "friend_recall"
	NoticeTypeNotify        NoticeType = "notify"
	NoticeTypeGroupCard     NoticeType = "group_card"
	NoticeTypeOfflineFile   NoticeType = "offline_file"
	NoticeTypeEssence       NoticeType = "essence"
)

type RequestType string

const (
	RequestTypeFriend RequestType = "friend"
	RequestTypeGroup  RequestType = "group"
)

type MetaEventType string

const (
	MetaEventTypeLifecycle MetaEventType = "lifecycle"
	MetaEventTypeHeartbeat MetaEventType = "heartbeat"
)

type Event struct {
	Time     int64    `json:"time"`
	SelfID   int64    `json:"self_id"`
	PostType PostType `json:"post_type"`
}

type MessageEvent struct {
	Event
	MessageType MessageType      `json:"message_type"`
	SubType     string           `json:"sub_type"`
	MessageID   int64            `json:"message_id"`
	UserID      int64            `json:"user_id"`
	Message     []MessageSegment `json:"message"`
	RawMessage  string           `json:"raw_message"`
	Font        int              `json:"font"`
	Sender      MessageSender    `json:"sender"`
}

type PrivateMessageEvent struct {
	MessageEvent
}

type GroupMessageEvent struct {
	MessageEvent
	GroupID   int64 `json:"group_id"`
	Anonymous *struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous,omitempty"`
}

type NoticeEvent struct {
	Event
	NoticeType NoticeType `json:"notice_type"`
}

type GroupUploadNoticeEvent struct {
	NoticeEvent
	GroupID int64      `json:"group_id"`
	UserID  int64      `json:"user_id"`
	File    *GroupFile `json:"file"`
}

type GroupAdminNoticeEvent struct {
	NoticeEvent
	SubType string `json:"sub_type"`
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
}

type GroupDecreaseNoticeEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}

type GroupIncreaseNoticeEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}

type GroupBanNoticeEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
	Duration   int64  `json:"duration"`
}

type FriendAddNoticeEvent struct {
	NoticeEvent
	UserID int64 `json:"user_id"`
}

type GroupRecallNoticeEvent struct {
	NoticeEvent
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	MessageID  int64 `json:"message_id"`
}

type FriendRecallNoticeEvent struct {
	NoticeEvent
	UserID    int64 `json:"user_id"`
	MessageID int64 `json:"message_id"`
}

type NotifyNoticeEvent struct {
	NoticeEvent
	SubType   string `json:"sub_type"`
	UserID    int64  `json:"user_id"`
	GroupID   int64  `json:"group_id,omitempty"`
	TargetID  int64  `json:"target_id,omitempty"`
	HonorType string `json:"honor_type,omitempty"`
}

type GroupCardNoticeEvent struct {
	NoticeEvent
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	CardNew string `json:"card_new"`
	CardOld string `json:"card_old"`
}

type OfflineFileNoticeEvent struct {
	NoticeEvent
	UserID int64 `json:"user_id"`
	File   struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
		URL  string `json:"url"`
	} `json:"file"`
}

type EssenceNoticeEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	SenderID   int64  `json:"sender_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int64  `json:"message_id"`
}

type RequestEvent struct {
	Event
	RequestType RequestType `json:"request_type"`
}

type FriendRequestEvent struct {
	RequestEvent
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}

type GroupRequestEvent struct {
	RequestEvent
	SubType string `json:"sub_type"`
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}

type MetaEvent struct {
	Event
	MetaEventType MetaEventType `json:"meta_event_type"`
}

type LifecycleMetaEvent struct {
	MetaEvent
	SubType string `json:"sub_type"`
}

type HeartbeatMetaEvent struct {
	MetaEvent
	Status struct {
		AppInitialized bool `json:"app_initialized"`
		AppEnabled     bool `json:"app_enabled"`
		AppGood        bool `json:"app_good"`
		Online         bool `json:"online"`
		Good           bool `json:"good"`
		Statistics     struct {
			PacketReceived  int64 `json:"packet_received"`
			PacketSent      int64 `json:"packet_sent"`
			PacketLost      int64 `json:"packet_lost"`
			MessageReceived int64 `json:"message_received"`
			MessageSent     int64 `json:"message_sent"`
			DisconnectTimes int64 `json:"disconnect_times"`
			LostTimes       int64 `json:"lost_times"`
			LastMessageTime int64 `json:"last_message_time"`
		} `json:"statistics"`
	} `json:"status"`
	Interval int64 `json:"interval"`
}
