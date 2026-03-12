package onebot

type GroupInfo struct {
	GroupID         int64  `json:"group_id"`
	GroupName       string `json:"group_name"`
	GroupUin        int64  `json:"group_uin"`
	AdminFlag       int    `json:"admin_flag"`
	GroupUinName    string `json:"group_uin_name"`
	LastMsgTime     int64  `json:"last_msg_time"`
	MaxMember       int    `json:"max_member"`
	MemberNum       int    `json:"member_num"`
	MemberCount     int    `json:"member_count"`
	GroupRemark     string `json:"group_remark"`
	GroupNotice     string `json:"group_notice"`
	GroupStatus     int    `json:"group_status"`
	GroupOwner      int64  `json:"group_owner"`
	GroupLevel      int    `json:"group_level"`
	GroupCreateTime int64  `json:"group_create_time"`
	GroupTotalCount int    `json:"group_total_count"`
	ShutUpTimestamp int64  `json:"shut_up_timestamp"`
	MyShutUp        int    `json:"my_shut_up"`
}

type GroupMemberInfo struct {
	GroupID         int64  `json:"group_id"`
	UserID          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int    `json:"age"`
	Area            string `json:"area"`
	JoinTime        int64  `json:"join_time"`
	LastSentTime    int64  `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
	ShutUpTimestamp int64  `json:"shut_up_timestamp"`
}

type GroupPokeRequest struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
}

type GetGroupInfoRequest struct {
	GroupID int64 `json:"group_id"`
	NoCache bool  `json:"no_cache,omitempty"`
}

type GetGroupMemberListRequest struct {
	GroupID int64 `json:"group_id"`
	NoCache bool  `json:"no_cache,omitempty"`
}

type GetGroupMemberInfoRequest struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache,omitempty"`
}

type SetGroupAddRequestRequest struct {
	Flag    string `json:"flag"`
	SubType string `json:"sub_type"`
	Approve bool   `json:"approve,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type SetGroupLeaveRequest struct {
	GroupID   int64 `json:"group_id"`
	IsDismiss bool  `json:"is_dismiss,omitempty"`
}

type SetGroupAdminRequest struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	Enable  bool  `json:"enable,omitempty"`
}

type SetGroupCardRequest struct {
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Card    string `json:"card,omitempty"`
}

type SetGroupBanRequest struct {
	GroupID  int64 `json:"group_id"`
	UserID   int64 `json:"user_id"`
	Duration int64 `json:"duration,omitempty"`
}

type SetGroupWholeBanRequest struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable,omitempty"`
}

type SetGroupNameRequest struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

type BatchDeleteGroupMemberRequest struct {
	GroupID int64   `json:"group_id"`
	UserIDs []int64 `json:"user_ids"`
	Enable  bool    `json:"enable,omitempty"`
}

type SetGroupKickRequest struct {
	GroupID          int64 `json:"group_id"`
	UserID           int64 `json:"user_id"`
	RejectAddRequest bool  `json:"reject_add_request,omitempty"`
}

type SetGroupSpecialTitleRequest struct {
	GroupID      int64  `json:"group_id"`
	UserID       int64  `json:"user_id"`
	SpecialTitle string `json:"special_title,omitempty"`
	Duration     int64  `json:"duration,omitempty"`
}

type GetGroupHonorInfoRequest struct {
	GroupID int64  `json:"group_id"`
	Type    string `json:"type,omitempty"`
}

type SetEssenceMsgRequest struct {
	MessageID int64 `json:"message_id"`
}

type DeleteEssenceMsgRequest struct {
	MessageID int64 `json:"message_id"`
}

type GetGroupAtAllRemainRequest struct {
	GroupID int64 `json:"group_id"`
}

type SendGroupNoticeRequest struct {
	GroupID int64  `json:"group_id"`
	Content string `json:"content"`
	Image   string `json:"image,omitempty"`
}

type GetGroupNoticeRequest struct {
	GroupID int64 `json:"group_id"`
}

type DeleteGroupNoticeRequest struct {
	GroupID  int64  `json:"group_id"`
	NoticeID string `json:"notice_id"`
}

type SendGroupSignRequest struct {
	GroupID int64 `json:"group_id"`
}

type SetGroupMsgMaskRequest struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupRemarkRequest struct {
	GroupID int64  `json:"group_id"`
	Remark  string `json:"remark"`
}

type GetGroupIgnoreAddRequestRequest struct {
	GroupID int64 `json:"group_id"`
}

type UploadGroupAlbumRequest struct {
	GroupID int64  `json:"group_id"`
	File    string `json:"file"`
	Name    string `json:"name,omitempty"`
}

type GetGroupAlbumListRequest struct {
	GroupID int64 `json:"group_id"`
}

type CreateGroupAlbumRequest struct {
	GroupID     int64  `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type DeleteGroupAlbumRequest struct {
	GroupID int64  `json:"group_id"`
	AlbumID string `json:"album_id"`
}

type GroupSystemMsg struct {
	InvitedRequests []GroupInvitedRequest `json:"invited_requests"`
	JoinRequests    []GroupJoinRequest    `json:"join_requests"`
}

type GroupInvitedRequest struct {
	RequestID   int64  `json:"request_id"`
	InvitorUin  int64  `json:"invitor_uin"`
	InvitorNick string `json:"invitor_nick"`
	GroupID     int64  `json:"group_id"`
	GroupName   string `json:"group_name"`
	Actioned    bool   `json:"actioned"`
	Action      string `json:"action"`
	Msg         string `json:"msg"`
	Flag        string `json:"flag"`
}

type GroupJoinRequest struct {
	RequestID     int64  `json:"request_id"`
	RequesterUin  int64  `json:"requester_uin"`
	RequesterNick string `json:"requester_nick"`
	Message       string `json:"message"`
	GroupID       int64  `json:"group_id"`
	GroupName     string `json:"group_name"`
	Actioned      bool   `json:"actioned"`
	Action        string `json:"action"`
	Flag          string `json:"flag"`
	Seq           int64  `json:"seq"`
}

type GroupShutInfo struct {
	UserID     int64  `json:"user_id"`
	Nickname   string `json:"nickname"`
	Duration   int64  `json:"duration"`
	ShutUpTime int64  `json:"shut_up_time"`
}

type EssenceMsg struct {
	SenderID     int64  `json:"sender_id"`
	SenderNick   string `json:"sender_nick"`
	MessageID    int64  `json:"message_id"`
	OperatorID   int64  `json:"operator_id"`
	OperatorNick string `json:"operator_nick"`
	MessageTime  int64  `json:"message_time"`
	OperatorTime int64  `json:"operator_time"`
}

type GroupAtAllRemain struct {
	CanAtAll                 bool `json:"can_at_all"`
	RemainAtAllCountForUin   int  `json:"remain_at_all_count_for_uin"`
	RemainAtAllCountForGroup int  `json:"remain_at_all_count_for_group"`
}

type GroupNotice struct {
	NoticeID    string `json:"notice_id"`
	SenderID    int64  `json:"sender_id"`
	PublishTime int64  `json:"publish_time"`
	Message     struct {
		Text   string `json:"text"`
		Images []struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			ID     string `json:"id"`
		} `json:"images"`
	} `json:"message"`
}

type GroupAlbum struct {
	AlbumID    string `json:"album_id"`
	AlbumName  string `json:"album_name"`
	CreateTime int64  `json:"create_time"`
	CreaterUin int64  `json:"creater_uin"`
	PicCount   int    `json:"pic_count"`
}

type GroupHonorInfo struct {
	GroupID          int64            `json:"group_id"`
	CurrentTalkative GroupHonorUser   `json:"current_talkative"`
	TalkativeList    []GroupHonorUser `json:"talkative_list"`
	PerformerList    []GroupHonorUser `json:"performer_list"`
	LegendList       []GroupHonorUser `json:"legend_list"`
	StrongNewbieList []GroupHonorUser `json:"strong_newbie_list"`
	EmotionList      []GroupHonorUser `json:"emotion_list"`
}

type GroupHonorUser struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	DayCount int    `json:"day_count,omitempty"`
	ID       int64  `json:"id,omitempty"`
	Desc     string `json:"description,omitempty"`
}
