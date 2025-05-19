package redis

import "context"

// InitRedis는 Redis 클라이언트를 초기화하는 함수
func InitRedis(ctx context.Context) error {
	if _, err := InitAuthRedis(ctx); err != nil {
		return err
	}

	return nil
}
