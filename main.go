package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "gindemo/docs"

	"gindemo/global"
	"gindemo/internal/model"
	"gindemo/internal/routers"
	"gindemo/pkg/logger"
	"gindemo/pkg/setting"
)

func init() {
	must(setupSetting())
	must(setupLogger())
	must(setupDBEngine())
}

// @title demo system
// @version 1.0
// @description gin framework
// @termsOfService https://www.baidu.com
func main() {
	router := routers.NewRouter()

	gin.SetMode(global.ServerSetting.RunMode)

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1M
	}
	global.Logger.Infof("Server starting... %#v", s)
	s.ListenAndServe()

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	fmt.Println(global.ServerSetting)
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	fmt.Println(global.AppSetting)
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	fmt.Println(global.DatabaseSetting)
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		},
		"",
		log.LstdFlags,
	).WithLevel(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
