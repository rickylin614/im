package errs

import "fmt"

const defaultErr = "00-000"

// group msg

var (
	commGroup                = Codes.Group("01")
	CommonUnknownError       = commGroup.Add("001", "未知错误")
	CommonNoData             = commGroup.Add("002", "查无资料")
	CommonRawSQLNotFound     = commGroup.Add("003", "找不到执行档")
	CommonServiceUnavailable = commGroup.Add("004", "系统维护中")
	CommonConfigureInvalid   = commGroup.Add("005", "设置参数错误")
	CommonParseError         = commGroup.Add("006", "解析失败")
)

var (
	requestGroup                  = Codes.Group("02")
	RequestParamInvalid           = requestGroup.Add("001", "请求参数错误")
	RequestParamParseFailed       = requestGroup.Add("002", "请求参数解析失败")
	RequestPageError              = requestGroup.Add("003", "请求的页数错误")
	RequestParseError             = requestGroup.Add("004", "解析失败")
	RequestParseTimeZoneError     = requestGroup.Add("005", "时区解析错误")
	RequestFrequentOperationError = requestGroup.Add("006", "频繁操作，请稍后再尝试")
	RequestTokenError             = requestGroup.Add("007", "token驗證失敗")
)

// ShowAllErrors 內部測試使用
func ShowAllErrors() {
	fmt.Println(CommonUnknownError.Error())       // 打印 CommonUnknownError 变量的错误消息
	fmt.Println(CommonNoData.Error())             // 打印 CommonNoData 变量的错误消息
	fmt.Println(CommonRawSQLNotFound.Error())     // 打印 CommonRawSQLNotFound 变量的错误消息
	fmt.Println(CommonServiceUnavailable.Error()) // 打印 CommonServiceUnavailable 变量的错误消息
	fmt.Println(CommonConfigureInvalid.Error())   // 打印 CommonConfigureInvalid 变量的错误消息
	fmt.Println(CommonParseError.Error())         // 打印 CommonParseError 变量的错误消息

	fmt.Println(RequestParamInvalid.Error())           // 打印 RequestParamInvalid 变量的错误消息
	fmt.Println(RequestParamParseFailed.Error())       // 打印 RequestParamParseFailed 变量的错误消息
	fmt.Println(RequestPageError.Error())              // 打印 RequestPageError 变量的错误消息
	fmt.Println(RequestParseError.Error())             // 打印 RequestParseError 变量的错误消息
	fmt.Println(RequestParseTimeZoneError.Error())     // 打印 RequestParseTimeZoneError 变量的错误消息
	fmt.Println(RequestFrequentOperationError.Error()) // 打印 RequestFrequentOperationError 变量的错误消息
}
