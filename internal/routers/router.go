package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gindemo/global"
	"gindemo/internal/middleware"
	"gindemo/internal/testmode"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Translations())

	swagerURL := ginSwagger.URL("http://127.0.0.1" + ":" + strconv.Itoa(global.ServerSetting.HttpPort) + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagerURL))
	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/tags", testmode.NewTag().Get)
	}

	return r
}
