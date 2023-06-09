package common

import (
	"github.com/gin-gonic/gin"

	"demo/internal/pkg/core"
	"demo/internal/pkg/log"
	v1 "demo/pkg/api/demo/v1"
)

// CommonController 是 common 模块在 Controller 层的实现，用来处理用户模块的请求.
type CommonController struct{}

// NewCommonController 创建一个 common controller.
func NewCommonController() *CommonController {
	return &CommonController{}
}

// Healthz
// @Summary    Heath check
// @Tags       Common
// @Accept     application/json
// @Produce    json
// @Success	   200		{object}	v1.HealthzResponse
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /healthz  [GET].
func (ctrl *CommonController) Healthz(c *gin.Context) {
	log.C(c).Infow("Healthz function called")

	data := &v1.HealthzResponse{Status: "ok"}

	core.WriteResponse(c, nil, data)
}
