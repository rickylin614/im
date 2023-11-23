package consts

// UserStatus 定義用戶狀態的枚舉或常量
type UserStatus int

const (
	UserStatusActive   UserStatus = iota // 正常狀態
	UserStatusBlocked                    // 被封鎖
	UserStatusInactive                   // 無效或未激活
)
