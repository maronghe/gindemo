package app

import (
	"github.com/gin-gonic/gin"

	"gindemo/global"
	"gindemo/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	switch {
	case pageSize <= 0:
		return global.AppSetting.DefaultPageSize
	case pageSize > global.AppSetting.MaxPageSize:
		return global.AppSetting.MaxPageSize
	default:
		return pageSize
	}
}
