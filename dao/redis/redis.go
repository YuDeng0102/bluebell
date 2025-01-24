package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"web-app/settings"
)

var rdb *redis.Client
var ctx = context.Background()

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port),
		Password: config.Password,
		DB:       config.Database,
		PoolSize: config.PoolSize,
	})

	_, err = rdb.Ping(ctx).Result()

	zap.L().Info("redis init success")
	return
}

func Close() {
	_ = rdb.Close()
}
