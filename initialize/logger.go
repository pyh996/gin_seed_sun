package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go_gin/global"
	"go_gin/utils"
)

// InitLogger 初始化Logger
func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}

	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	global.Lg = logg
}
