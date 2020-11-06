package global

import (
	"github.com/jinzhu/gorm"
	"github.com/lee820/ServerTemplate/pkg/logger"
	"github.com/lee820/ServerTemplate/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
