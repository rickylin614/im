package errs

import (
	"fmt"
	"net/http"
)

const defaultErr = "00-000"

// error from server
var (
	commGroup                = Codes.Group("01")
	CommonUnknownError       = commGroup.Add("未知错误", http.StatusInternalServerError)
	CommonServiceUnavailable = commGroup.Add("系统维护中", http.StatusServiceUnavailable)
	CommonConfigureInvalid   = commGroup.Add("设置参数错误", http.StatusBadRequest)
	CommonParseError         = commGroup.Add("解析失败", http.StatusBadRequest)
)

// error from client
var (
	requestGroup                  = Codes.Group("02")
	RequestParamInvalid           = requestGroup.Add("请求参数错误", http.StatusBadRequest)
	RequestParamParseFailed       = requestGroup.Add("请求参数解析失败", http.StatusBadRequest)
	RequestPageError              = requestGroup.Add("请求的页数错误", http.StatusBadRequest)
	RequestParseError             = requestGroup.Add("解析失败", http.StatusBadRequest)
	RequestParseTimeZoneError     = requestGroup.Add("时区解析错误", http.StatusBadRequest)
	RequestFrequentOperationError = requestGroup.Add("频繁操作，请稍后再尝试", http.StatusTooManyRequests)
	RequestNoData                 = requestGroup.Add("查无资料", http.StatusNotFound)
	RequestRawSQLNotFound         = requestGroup.Add("找不到执行档", http.StatusNotFound)
)

// 驗證錯誤
var (
	validGroup        = Codes.Group("03")
	RequestTokenError = validGroup.Add("登入失效，請重新登入", http.StatusUnauthorized)
	LoginCommonError  = validGroup.Add("使用者名稱或密碼無效", http.StatusUnauthorized)
	LoginLockedError  = validGroup.Add("使用者已被封鎖，請聯繫管理員", http.StatusForbidden)
)

// ShowAllErrors 內部測試使用
func ShowAllErrors() {
	group := []*GroupError{
		commGroup,
		requestGroup,
		validGroup,
	}

	for _, v := range group {
		for _, v2 := range v.ListCodeNMsg() {
			fmt.Println(v2)
		}
	}
}
