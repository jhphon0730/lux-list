package redis

import (
	"lux-list/internal/config"
	"lux-list/pkg/utils"

	"github.com/go-redis/redis/v8"
)

var (
	// auth_redis는 Redis 클라이언트 인스턴스
	auth_redis *redis.Client
)

// InitAuthRedis는 Redis 클라이언트를 초기화하고 반환하는 함수
func InitAuthRedis() *redis.Client {
	config := config.GetConfig()
	auth_redis = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       utils.InterfaceToInt(config.Redis.AuthDB),
	})

	return auth_redis
}
