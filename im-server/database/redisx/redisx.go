package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/laixhe/goutil/zaplog"

	"im-server/config"
)

var client redis.Cmdable

// Ping 测试
func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := client.Ping(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}

// DB -
func DB() redis.Cmdable {
	return client
}

func init() {
	nodes := config.RedisNodes()
	zaplog.Debug("▶▶▶Redis初始化...")
	if len(nodes) == 1 {
		// 单机
		client = redis.NewClient(&redis.Options{
			Addr:     nodes[0],
			Password: config.RedisPassword(),
			DB:       config.RedisDBNum(),
		})
	} else {
		// 分布式集群
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    nodes,
			Password: config.RedisPassword(),
		})
	}
	err := Ping()
	if err != nil {
		panic(err)
	}
	zaplog.Debug("▶▶▶Redis初始化完毕...")
}
