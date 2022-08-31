package global

import (
	"github.com/iamzhiyudong/xigua-blog/pkg/logger"
	"github.com/iamzhiyudong/xigua-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger

	JWTSetting *setting.JWTSettingS
)
