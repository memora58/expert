package init

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"task/common/global"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 这个对象如何在其他文件中使用 —— 全局变量
	if err := viper.Unmarshal(global.Config); err != nil {
		zap.S().Panic(err)
	}
}
