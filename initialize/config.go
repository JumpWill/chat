package initialize

import (
	"chat/global"

	"go.uber.org/zap"

	"chat/schema"
	"chat/utils"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	zap.L().Info("InitConfig", zap.String("mes", "config init"))

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// viper.SetConfigFile("../conf/conf.toml")
	viper.SetConfigFile("/Users/will/doc/code/golang/chat/conf/conf.toml")
	// 然后从环境变量读取
	viper.AutomaticEnv()
	// 将.替换为_
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	utils.Error(err, "read config failed")

	serverConfig := schema.ServerConfig{}

	err = viper.Unmarshal(&serverConfig)
	//给serverConfig初始值
	utils.Error(err, "init value failed")

	global.Settings = serverConfig
	zap.L().Info("InitConfig", zap.String("mes", "config init ok"))

}
