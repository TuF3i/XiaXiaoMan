package onebot

type FriendInfo struct {
	UserID        int64  `json:"user_id"`
	Nickname      string `json:"nickname"`
	Remark        string `json:"remark"`
	Sex           string `json:"sex"`
	BirthdayYear  int    `json:"birthday_year"`
	BirthdayMonth int    `json:"birthday_month"`
	BirthdayDay   int    `json:"birthday_day"`
	Age           int    `json:"age"`
	Qid           string `json:"qid"`
	LongNick      string `json:"long_nick"`
	Level         int    `json:"level,omitempty"`
	LongNick2     string `json:"longNick,omitempty"`
	Email         string `json:"eMail,omitempty"`
	Uid           string `json:"uid,omitempty"`
	CategoryID    int    `json:"categoryId,omitempty"`
	RichTime      int    `json:"richTime,omitempty"`
}

type FriendCategory struct {
	CategoryID      int          `json:"categoryId"`
	CategorySortID  int          `json:"categorySortId"`
	CategoryName    string       `json:"categoryName"`
	CategoryMbCount int          `json:"categoryMbCount"`
	OnlineCount     int          `json:"onlineCount"`
	BuddyList       []FriendInfo `json:"buddyList"`
}

type StrangerInfo struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
	Qid         string `json:"qid"`
	Level       int    `json:"level"`
	LoginDays   int    `json:"login_days"`
	VipLevel    int    `json:"vip_level"`
	LongNick    string `json:"long_nick"`
	City        string `json:"city"`
	Sign        string `json:"sign"`
	QqLevel     int    `json:"qq_level"`
	UserExtInfo string `json:"user_ext_info"`
	UserAddTime int64  `json:"user_add_time"`
	IsFriend    bool   `json:"is_friend"`
	NicknameOld string `json:"nickname_old"`
}

type ProfileLike struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Source      int    `json:"source"`
	Time        int64  `json:"time"`
	EmojiID     string `json:"emoji_id"`
	SendAvatar  string `json:"send_avatar"`
	EmojiPackID int64  `json:"emoji_pack_id"`
	EmojiURL    string `json:"emoji_url"`
}

type RobotUinRange struct {
	MaxRobotUin int64 `json:"max_robot_uin"`
	MinRobotUin int64 `json:"min_robot_uin"`
}

type DoubtFriendAddRequest struct {
	RequestID    int64  `json:"request_id"`
	RequesterUin int64  `json:"requester_uin"`
	Message      string `json:"message"`
	Flag         string `json:"flag"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Nickname     string `json:"nickname"`
	AddSource    string `json:"add_source"`
	AddWayName   string `json:"add_way_name"`
}

type QQAvatar struct {
	FaceID      int    `json:"face_id"`
	FaceUrl     string `json:"face_url"`
	FaceOld     string `json:"face_old"`
	FaceFileMd5 string `json:"face_file_md5"`
}
