package initialize

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"go_gin/global"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 链接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
	}
}
