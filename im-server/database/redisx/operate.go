package redisx

import (
	"context"
	"time"
)

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

// Get Redis GET 命令用于获取指定 key 的值
func Get(key string) (string, error) {
	cmd := client.Get(context.Background(), key)
	return cmd.Val(), cmd.Err()
}

// Set Redis SET 命令用于设置给定 key 的值，如果 key 已经存储其他值，SET 就覆写旧值，且无视类型
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(context.Background(), key, value, expiration).Err()
}

// Incr Redis INCR 将 key 中储存的数字值增一
// 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作
func Incr(key string) (int64, error) {
	cmd := client.Incr(context.Background(), key)
	return cmd.Val(), cmd.Err()
}
