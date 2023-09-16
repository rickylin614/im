package middleware

import (
	"im/internal/consts"
	"im/internal/util/ctxs"
	"im/internal/util/errs"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	in digIn
}

func (m authMiddleware) IsLogin(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if len(token) == 0 {
		ctxs.SetError(ctx, errs.RequestTokenError)
		ctx.Abort()
		return
	}

	user, err := m.in.Service.UsersSrv.GetByToken(ctx, token)
	if err != nil || user == nil || user.Status != consts.UserStatusActive {
		ctxs.SetError(ctx, errs.RequestTokenError)
		ctx.Abort()
		return
	}
	ctxs.SetUserInfo(ctx, user)
	ctx.Next()
}
