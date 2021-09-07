package redistest

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func RedisTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "[2400:8902::f03c:92ff:fe51:5871]:6379",
		Username: "milittle",
		Password: "abc",
		DB:       0,
	})

	//fmt.Println(rdb.Ping(ctx))

	err := rdb.Set(ctx, "key", "value", time.Second*500).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
