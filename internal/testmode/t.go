package testmode

import (
	"github.com/gin-gonic/gin"

	"gindemo/pkg/app"
	"gindemo/pkg/errorcode"
)

type Tag struct {
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)
	return
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default 1
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} testmode.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)
	return
}
