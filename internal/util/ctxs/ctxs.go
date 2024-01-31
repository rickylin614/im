package ctxs

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"

	"im/internal/models/po"
	"im/internal/models/response"
	"im/internal/pkg/consts"
	"im/internal/pkg/consts/enums"
	"im/internal/util/errs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IMessage interface {
	GetMessage() string
}

type IError interface {
	GetCode() string
}

type IStatusCode interface {
	GetStatusCode() int
}

func Background() *gin.Context {
	return &gin.Context{}
}

func SetError(ctx *gin.Context, err error) {
	code, msg, data, statusCode := ParseError(err)
	response := response.APIResponse[any]{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}

// SetResp 設定一般回傳格式
func SetResp(ctx *gin.Context, data any) {
	msg := ""
	if v, ok := data.(IMessage); ok {
		msg = v.GetMessage()
	}
	response := response.APIResponse[any]{
		Code:    "0",
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func SetSuccessResp(ctx *gin.Context) {
	response := response.APIResponse[any]{
		Code:    "0",
		Message: "操作成功",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}

func SetUserInfo(ctx *gin.Context, user *po.Users) {
	ctx.Set(consts.UserInfo, user)
}

func GetUserInfo(ctx *gin.Context) (user *po.Users) {
	data, ok := ctx.Get(consts.UserInfo)
	if !ok {
		return &po.Users{}
		//panic(errors.New("not Login Func Use UserInfo"))
	}
	user, ok = data.(*po.Users)
	if !ok {
		return &po.Users{}
		//panic(errors.New("not Login Func Use UserInfo"))
	}
	return
}

func GetDeviceID(ctx *gin.Context) string {
	deviceID := ctx.GetHeader("X-Device-ID")
	if deviceID != "" {
		return deviceID
	}

	// 如果header中沒有X-Device-ID，嘗試從User-Agent解析
	userAgent := ctx.GetHeader("User-Agent")
	return parseDeviceIDFromUserAgent(userAgent)
}

// parseDeviceIDFromUserAgent 簡易從userAgent中分析設備
func parseDeviceIDFromUserAgent(userAgent string) string {
	if strings.Contains(userAgent, "iPhone") {
		return "iPhone"
	} else if strings.Contains(userAgent, "Android") {
		return "Android"
	}
	return "Default"
}

func ParseError(err error) (code string, msg string, data any, statusCode int) {
	statusCode = http.StatusBadRequest
	code = errs.CommonUnknownError.GetCode()
	msg = errs.CommonUnknownError.GetMessage()
	if data, ok := ParseBindingErrMsg(err); ok {
		return errs.RequestParamInvalid.GetCode(), errs.RequestParamInvalid.GetMessage(), data, statusCode
	}
	if data, ok := ParseMySQLError(err); ok {
		return errs.CommonSQLExecutionError.GetCode(), errs.CommonSQLExecutionError.GetMessage(), data, statusCode
	}

	if v, ok := err.(IMessage); ok {
		msg = v.GetMessage()
	}
	if v, ok := err.(IError); ok {
		code = v.GetCode()
	}
	if v, ok := err.(IStatusCode); ok {
		statusCode = v.GetStatusCode()
	}

	return code, msg, err.Error(), statusCode
}

// ParseBindingErrMsg
// 轉換binding錯誤
func ParseBindingErrMsg(err error) ([]string, bool) {
	var ValidationErrors validator.ValidationErrors
	if !errors.As(err, &ValidationErrors) {
		return nil, false
	}

	var errorMessages []string
	for _, fieldErr := range ValidationErrors {
		switch fieldErr.Tag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 是必填字段", fieldErr.Field()))
		case "alphanum":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 只能包含字母和数字字符", fieldErr.Field()))
		case "min":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 长度必须至少为 %s 个字符", fieldErr.Field(), fieldErr.Param()))
		case "max":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 长度必须不可多余 %s 个字符", fieldErr.Field(), fieldErr.Param()))
		case "email":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不是有效的电子邮件地址", fieldErr.Field()))
		case "numeric":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 只能包含数字字符", fieldErr.Field()))
		case "len":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须为 %s 个字符", fieldErr.Field(), fieldErr.Param()))
		case "eq":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须等于 %s", fieldErr.Field(), fieldErr.Param()))
		case "eq_ignore_case":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须等于 %s（不区分大小写）", fieldErr.Field(), fieldErr.Param()))
		case "gt":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须大于 %s", fieldErr.Field(), fieldErr.Param()))
		case "gte":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须大于或等于 %s", fieldErr.Field(), fieldErr.Param()))
		case "lt":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须小于 %s", fieldErr.Field(), fieldErr.Param()))
		case "lte":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须小于或等于 %s", fieldErr.Field(), fieldErr.Param()))
		case "ne":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不能等于 %s", fieldErr.Field(), fieldErr.Param()))
		case "ne_ignore_case":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不能等于 %s（不区分大小写）", fieldErr.Field(), fieldErr.Param()))
		case "filepath":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不是有效的文件路径", fieldErr.Field()))
		case "image":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不是有效的图像", fieldErr.Field()))
		case "timezone":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不是有效的时区", fieldErr.Field()))
		case "alphaunicode":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 只能包含 Unicode 字母", fieldErr.Field()))
		case "uuid":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 不是有效的 UUID", fieldErr.Field()))
		case "oneof":
			errorMessages = append(errorMessages, fmt.Sprintf("%s 必须是以下选项之一: %s", fieldErr.Field(), fieldErr.Param()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s 无效参数", fieldErr))
		}
	}

	return errorMessages, true
}

func ParseMySQLError(err error) ([]string, bool) {
	var mySQLError *mysql.MySQLError
	if !errors.As(err, &mySQLError) {
		return nil, false
	}

	var errorMessages []string
	switch mySQLError.Number {
	case 1045:
		errorMessages = append(errorMessages, "无效的用户名或密码")
	case 1049:
		errorMessages = append(errorMessages, "未知的数据库")
	case 1054:
		errorMessages = append(errorMessages, "未知的列")
	case 1062:
		// 正则表达式来提取重复键的名称
		re := regexp.MustCompile(`for key '(.+)'`)
		matches := re.FindStringSubmatch(mySQLError.Message)
		if len(matches) > 1 {
			errorMessages = append(errorMessages, fmt.Sprintf("重复的条目 '%s'", matches[0]))
		} else {
			errorMessages = append(errorMessages, "重复的条目")
		}
	case 1064:
		errorMessages = append(errorMessages, "SQL 语法错误")
	case 1136:
		errorMessages = append(errorMessages, "列数不匹配")
	case 1146:
		// 正则表达式来提取表名
		re := regexp.MustCompile(`Table '(.+)' doesn't exist`)
		matches := re.FindStringSubmatch(mySQLError.Message)
		if len(matches) > 1 {
			errorMessages = append(errorMessages, fmt.Sprintf("未知的表 '%s'", matches[0]))
		} else {
			errorMessages = append(errorMessages, "未知的表")
		}
	case 1217:
		errorMessages = append(errorMessages, "存在外键约束")
	case 1451:
		errorMessages = append(errorMessages, "无法删除或更新父行")
	case 1452:
		errorMessages = append(errorMessages, "不能添加或更新子行")
	case 2002, 2003:
		errorMessages = append(errorMessages, "无法连接到 MySQL 服务器")
	case 2006:
		errorMessages = append(errorMessages, "MySQL 服务器已经断开连接")
	case 2013:
		errorMessages = append(errorMessages, "在与 MySQL 服务器通信时丢失连接")
	case 2026:
		errorMessages = append(errorMessages, "SSL 连接错误")
	case 2055:
		errorMessages = append(errorMessages, "失去与 MySQL 服务器的连接")
	default:
		errorMessages = append(errorMessages, fmt.Sprintf("MySQL错误 [%d]: %s", mySQLError.Number, mySQLError.Message))
	}

	return errorMessages, true
}

// IGetToken 解藕用interface
type IGetToken interface {
	GetByToken(ctx *gin.Context, token string) (user *po.Users, err error)
}

func CheckLoginByParam(ctx *gin.Context, userSrv IGetToken) error {
	token := ctx.Query(consts.Authorization_Header)
	user, err := userSrv.GetByToken(ctx, token)
	if err != nil || user == nil || user.Status != enums.UserStatusActive {
		return errs.RequestTokenError
	}
	SetUserInfo(ctx, user)
	return nil
}
