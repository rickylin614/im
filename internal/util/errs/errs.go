package errs

import (
	"fmt"
	"net/http"
)

const defaultErr = "00-000"

var (
	commGroup                = Codes.Group("01")                                        // error from server
	CommonUnknownError       = commGroup.Add("未知错误", http.StatusInternalServerError)    // 未知错误 (HTTP 500)
	CommonServiceUnavailable = commGroup.Add("系统维护中", http.StatusServiceUnavailable)    // 系统维护中 (HTTP 503)
	CommonConfigureInvalid   = commGroup.Add("设置参数错误", http.StatusBadRequest)           // 设置参数错误 (HTTP 400)
	CommonParseError         = commGroup.Add("解析失败", http.StatusBadRequest)             // 解析失败 (HTTP 400)
	CommonSQLExecutionError  = commGroup.Add("服务器执行错误", http.StatusInternalServerError) // 服务器执行错误 (HTTP 500)
)

var (
	requestGroup                  = Codes.Group("02")                                           // error from client
	RequestParamInvalid           = requestGroup.Add("请求参数错误", http.StatusBadRequest)           // 请求参数错误 (HTTP 400)
	RequestParamParseFailed       = requestGroup.Add("请求参数解析失败", http.StatusBadRequest)         // 请求参数解析失败 (HTTP 400)
	RequestPageError              = requestGroup.Add("请求的页数错误", http.StatusBadRequest)          // 请求的页数错误 (HTTP 400)
	RequestParseError             = requestGroup.Add("解析失败", http.StatusBadRequest)             // 解析失败 (HTTP 400)
	RequestParseTimeZoneError     = requestGroup.Add("时区解析错误", http.StatusBadRequest)           // 时区解析错误 (HTTP 400)
	RequestFrequentOperationError = requestGroup.Add("频繁操作，请稍后再尝试", http.StatusTooManyRequests) // 频繁操作，请稍后再尝试 (HTTP 429)
	RequestNoData                 = requestGroup.Add("查无资料", http.StatusNotFound)               // 查无资料 (HTTP 404)
	RequestRawSQLNotFound         = requestGroup.Add("找不到执行档", http.StatusNotFound)             // 找不到执行档 (HTTP 404)
	RequestDuplicate              = requestGroup.Add("请求重复", http.StatusConflict)               // 请求重复 (HTTP 409)
	RequestInvalidUser            = requestGroup.Add("无效的用户", http.StatusBadRequest)            // 无效的用户 (HTTP 400)
	RequestInvalidID              = requestGroup.Add("无效的ID", http.StatusBadRequest)            // 无效的ID (HTTP 400)
	RequestInvalidPermission      = requestGroup.Add("权限不足", http.StatusBadRequest)             // 权限不足 (HTTP 400)
)

var (
	loginGroup        = Codes.Group("03")                                      // 登入相關
	RequestTokenError = loginGroup.Add("登入失效，請重新登入", http.StatusUnauthorized)  // 登入失效，請重新登入 (HTTP 401)
	LoginCommonError  = loginGroup.Add("使用者名稱或密碼無效", http.StatusUnauthorized)  // 使用者名稱或密碼無效 (HTTP 401)
	LoginLockedError  = loginGroup.Add("使用者已被封鎖，請聯繫管理員", http.StatusForbidden) // 使用者已被封鎖，請聯繫管理員 (HTTP 403)
)

var (
	GroupGroup            = Codes.Group("03")                                 // 登入相關
	GroupMemberExistError = loginGroup.Add("對象已是群組成員", http.StatusBadRequest) // 對象已是群組成員 (HTTP 400)
)

var (
	businessGroup          = Codes.Group("04")                              // 業務相關
	BusinessFriendshipHint = businessGroup.Add("您已经是该用户的好友", http.StatusOK) // 您已经是该用户的好友 (HTTP 200)
)

var (
	WebSocketGroup               = Codes.Group("05")                                              // 長連線相關
	WebSocketMaxConnectionsError = WebSocketGroup.Add("已達到最大連線數量", http.StatusServiceUnavailable) // 達到最大連線數量 (HTTP 503)
	WebSocketConnectionTimeout   = WebSocketGroup.Add("連線超時", http.StatusRequestTimeout)          // 連線超時 (HTTP 408)
	WebSocketInvalidDataFormat   = WebSocketGroup.Add("資料格式錯誤", http.StatusBadRequest)            // 資料格式錯誤 (HTTP 400)
	WebSocketUnauthorizedAccess  = WebSocketGroup.Add("未授權的存取", http.StatusUnauthorized)          // 未授權的存取 (HTTP 401)
	WebSocketConnectionClosed    = WebSocketGroup.Add("連線已被關閉", http.StatusGone)                  // 連線已被關閉 (HTTP 410)
	WebSocketInternalServerError = WebSocketGroup.Add("服務器內部錯誤", http.StatusInternalServerError)  // 服務器內部錯誤 (HTTP 500)
)

// ShowAllErrors 內部測試使用
func ShowAllErrors() {
	group := []*GroupError{
		commGroup,
		requestGroup,
		loginGroup,
	}

	for _, v := range group {
		for _, v2 := range v.ListCodeNMsg() {
			fmt.Println(v2)
		}
	}
}
