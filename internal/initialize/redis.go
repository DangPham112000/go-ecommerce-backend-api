package initialize

import (
	"context"
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port), // "localhost:6379",
		Password: r.Password,                           // no password set
		DB:       r.Database,                           // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis initialization error")
	}

	fmt.Println("Redis is running")

	global.Rbd = rdb

}
