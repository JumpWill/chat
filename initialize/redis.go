package initialize

import (
	"chat/constant"
	"chat/global"
	"chat/utils"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() {

	utils.Logger(constant.Info, map[string]string{"mes": "redis init"})
	addr := fmt.Sprintf("%s:%d", global.Settings.Redis.Host, global.Settings.Redis.Port)
	Redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.Settings.Redis.Password,
		DB:       global.Settings.Redis.DB,
	})
	_, err := Redis.Ping().Result()
	utils.Error(err, "failed to connect redis")
	utils.Logger(constant.Info, map[string]string{"mes": "redis init ok"})

}
