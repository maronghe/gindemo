package testmode

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gindemo/global"
	"gindemo/pkg/app"
	"gindemo/pkg/errcode"
)

type Tag struct {
}

// used for swagger...
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default 1
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} testmode.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

	params := struct {
		Name  string `form:"name" binding="max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if errs != nil {
		global.Logger.Warnf("BindAndValid err %+v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(errcode.Success.WithDetails(fmt.Sprintf("valid is %t", valid), "669"))
	return
}
