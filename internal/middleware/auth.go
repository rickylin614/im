package middleware

import (
	"im/internal/pkg/consts/enums"
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
		token = ctx.GetHeader("authorization")
		if len(token) == 0 {
			ctxs.SetError(ctx, errs.RequestTokenError)
			ctx.Abort()
			return
		}
	}

	user, err := m.in.Service.UsersSrv.GetByToken(ctx, token)
	if err != nil || user == nil || user.Status != enums.UserStatusActive {
		ctxs.SetError(ctx, errs.RequestTokenError)
		ctx.Abort()
		return
	}
	ctxs.SetUserInfo(ctx, user)
	ctx.Next()
}
