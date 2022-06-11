package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisClient *redis.Client

func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "121.5.132.57:6700",
		Password: "123456", // no password set
		DB:       1,        // use DB1
	})
	//测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	RedisClient = rdb
}
