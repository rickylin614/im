package resp

type UsersGet struct {
	ID          string `json:"id"`           // uid
	Username    string `json:"username"`     // 用戶名稱
	Nickname    string `json:"nickname"`     // 用戶暱稱
	Email       string `json:"email"`        // 電子郵件地址
	Password    string `json:"password"`     // 密碼
	PhoneNumber string `json:"phone_number"` // 手機號碼
}

type UsersGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []UsersGet   `json:"data"`
}

type UsersLogin struct {
	Token string `json:"token"` // 登入Token
}
