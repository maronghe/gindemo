package routers

import (
	"gindemo/global"
	"gindemo/internal/service"
	"gindemo/pkg/app"
	"gindemo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	params := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %+v", errs)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&params)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %+v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExit)
		return
	}
	token, err := app.GenerateToken(params.AppKey, params.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %+v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
