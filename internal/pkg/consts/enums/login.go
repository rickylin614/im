package enums

// 定義用戶狀態的枚舉或常量
type LoginState int

// 登入狀態
const (
	LoginStateSuccess LoginState = iota // 登入成功
	LoginStateFailed                    // 登入失敗
	LoginStateBlocked                   // 帳號被封鎖
)
