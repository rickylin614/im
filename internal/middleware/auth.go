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
		return
	}

	// 存取redsi 若已經有資料且可轉models.User則Pass
	// redisdb := utils.GetRedisDb()
	// cmd := redisdb.Get(constant.LoginKey + data)
	// if cmd.Err() != nil || cmd.Val() == "" {
	// 	fmt.Printf("err: %v , value %v\n", cmd.Err(), cmd.Val())
	// 	ctx.JSON(http.StatusOK, out)
	// 	ctx.Abort()
	// 	return
	// } else {
	// 	redisdb.Expire(constant.LoginKey, time.Hour*3)
	// 	var user models.User
	// 	err := json.Unmarshal([]byte(cmd.Val()), &user)
	// 	if err != nil {
	// 		zapLog.ErrorW("login check err!", err)
	// 		ctx.Abort()
	// 		return
	// 	}
	// 	ctx.Set("user", user)
	// }
	ctx.Next()
}
