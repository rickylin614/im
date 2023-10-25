package ctxs

import (
	"errors"
	"fmt"
	"net/http"

	"im/internal/models"
	"im/internal/models/resp"
	"im/internal/pkg/consts"
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

func SetError(ctx *gin.Context, err error) {
	code, msg, data, statusCode := ParseError(err)
	response := resp.APIResponse[any]{
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
	response := resp.APIResponse[any]{
		Code:    "0",
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func SetSuccessResp(ctx *gin.Context) {
	response := resp.APIResponse[any]{
		Code:    "0",
		Message: "操作成功",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}

func SetUserInfo(ctx *gin.Context, user *models.Users) {
	ctx.Set(consts.UserInfo, user)
}

func GetUserInfo(ctx *gin.Context) (user *models.Users) {
	data, ok := ctx.Get(consts.UserInfo)
	if !ok {
		panic(errors.New("Not Login Func Use UserInfo"))
	}
	user, ok = data.(*models.Users)
	if !ok {
		panic(errors.New("Not Login Func Use UserInfo"))
	}
	return
}

func ParseError(err error) (code string, msg string, data any, statusCode int) {
	statusCode = http.StatusBadRequest
	code = errs.CommonUnknownError.GetCode()
	msg = errs.CommonUnknownError.GetMessage()
	if data, ok := ParseBindingErrMsg(err); ok {
		return errs.RequestParamInvalid.GetCode(), errs.RequestParamInvalid.GetMessage(), data, statusCode
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
	ValidationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
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
