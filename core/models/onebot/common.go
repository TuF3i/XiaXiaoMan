package onebot

type APIResponse struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Wording string      `json:"wording"`
	Echo    string      `json:"echo,omitempty"`
}

type APIRequest struct {
	Action string      `json:"action"`
	Params interface{} `json:"params,omitempty"`
	Echo   string      `json:"echo,omitempty"`
}

type SendLikeRequest struct {
	UserID int64 `json:"user_id"`
	Times  int   `json:"times,omitempty"`
}

type GetStrangerInfoRequest struct {
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache,omitempty"`
}

type SetFriendAddRequestRequest struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve,omitempty"`
	Remark  string `json:"remark,omitempty"`
}

type SetFriendRemarkRequest struct {
	UserID int64  `json:"user_id"`
	Remark string `json:"remark"`
}

type DeleteFriendRequest struct {
	UserID int64 `json:"user_id"`
}

type FriendPokeRequest struct {
	UserID int64 `json:"user_id"`
}

type SetQQAvatarRequest struct {
	File string `json:"file"`
}

type SetQQProfileRequest struct {
	Nickname     string `json:"nickname,omitempty"`
	Company      string `json:"company,omitempty"`
	Email        string `json:"email,omitempty"`
	College      string `json:"college,omitempty"`
	PersonalNote string `json:"personal_note,omitempty"`
}

type SetFriendCategoryRequest struct {
	UserID     int64  `json:"user_id"`
	CategoryID int    `json:"category_id,omitempty"`
	Category   string `json:"category,omitempty"`
}

type GetDoubtFriendsAddRequestRequest struct {
	NoCache bool `json:"no_cache,omitempty"`
}

type SetDoubtFriendsAddRequestRequest struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve"`
}
