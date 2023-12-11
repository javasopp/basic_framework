package redis

import (
	"context"
	"fmt"
	"gin-basic/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func Init() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", config.Config.Redis.Host, config.Config.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Database,
	})

	// 测试Redis连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
