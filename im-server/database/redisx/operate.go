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
