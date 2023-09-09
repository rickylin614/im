package ctxs

import (
	"im/internal/models/resp"
	"im/internal/util/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IMessage interface {
	GetMessage() string
}

type IError interface {
	GetCode() string
}

func SetError(ctx *gin.Context, err error) {
	msg := ""
	code := ""
	if v, ok := err.(IMessage); ok {
		msg = v.GetMessage()
	}
	if v, ok := err.(IError); ok {
		code = v.GetCode()
	} else {
		code = errs.CommonUnknownError.GetCode()
		msg = errs.CommonUnknownError.GetMessage()
	}
	response := resp.APIResponse[any]{
		Code:    code,
		Message: msg,
		Data:    err.Error(),
	}
	ctx.JSON(http.StatusOK, response)
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
