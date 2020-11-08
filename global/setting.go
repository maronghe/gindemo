package global

import (
	"github.com/jinzhu/gorm"

	"gindemo/pkg/logger"
	"gindemo/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS

	DBEngine *gorm.DB

	Logger *logger.Logger
)
