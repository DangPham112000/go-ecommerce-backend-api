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
	s := fmt.Sprintf("%s:%v", r.Host, r.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     s,          // "localhost:6379",
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis initialization error")
	}

	global.Logger.Info("Redis initialization success!!!")
	// global.Logger.Info(fmt.Sprintf("Redis initialization success!!! s=%v", s))

	global.Rbd = rdb

}
