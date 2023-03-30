package redis

import (
	"context"
	"gateway/common/global"
	"github.com/go-redis/redis/v8"
)

func RedisConnectFactory(db int) (*redis.Client, error) {
	//if global.RedisMapPool[db] != nil {
	//	return global.RedisMapPool[db], nil
	//}
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     redisConfig.Addr,
	//	Password: redisConfig.Password, // no password set
	//	DB:       db,                   // use default DB
	//})
	//
	//_, err := rdb.Ping(context.Background()).Result()
	//if err != nil {
	//	return nil, err
	//}
	//global.RedisMapPool[db] = rdb
	//return global.RedisMapPool[db], nil

	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password, // no password set
		DB:       db,                           // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
