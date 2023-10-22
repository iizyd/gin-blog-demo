package global

import (
	"github.com/iizyd/xigua-blog/pkg/logger"
	"github.com/iizyd/xigua-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger

	JWTSetting *setting.JWTSettingS
)
