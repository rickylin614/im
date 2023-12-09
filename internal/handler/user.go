package handler

import (
	request "im/internal/models/req"
	response "im/internal/models/resp"
	"im/internal/util/ctxs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type userHandler struct {
	in digIn
}

// Get
// @Summary Get
// @Tags user
// @Param request body request.UserGet true "param"
// @Success 200 {object} response.APIResponse[response.UserGet]
// @Router /