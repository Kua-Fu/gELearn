package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func getRdb() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "abc123", // no password set
		DB:       0,        // use default DB
	})
	return rdb
}

func TestRedis(t *testing.T) {
	rdb := getRdb()
	err := rdb.Set(ctx, "key", "value", 0).Err()
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
	// Output: key value
	// key2 does not exist
}

func TestRedisMap(t *testing.T) {
	rdb := getRdb()
	ip := "10.23.190.38"
	oriDstHost := "yz_test_1"
	// 查看是否存在
	now := time.Now().Unix()

	dstHost, err := rdb.HGet(ctx, "network_dst_host", ip).Result()
	// 1) 不存在field
	if err != nil {
		fmt.Println("--err--", dstHost, err)
		newDstHost := fmt.Sprintf("%d-%s", now, oriDstHost)
		fmt.Println("--new--", newDstHost)
		rdb.HSet(ctx, "network_dst_host", ip, newDstHost).Result()
	}
	// 2) 存在 field，判断是否过期
	s := strings.Split(dstHost, "-")
	if len(s) > 0 {
		setTime, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return
		}
		// 超过有效期
		if now-setTime > 60*60*24 {
			rdb.HDel(ctx, "network_dst_host", ip).Result()
		}
	}

}
