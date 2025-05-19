package redis

import (
	"lux-list/internal/config"
	"lux-list/pkg/utils"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	// auth_redis는 Redis 클라이언트 인스턴스
	auth_redis *redis.Client
	// auth_redis_once는 Redis 클라이언트 초기화를 위한 sync.Once 인스턴스
	auth_redis_once sync.Once
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

// GetAuthRedis는 Redis 클라이언트를 반환하는 함수
func GetAuthRedis() *redis.Client {
	// auth_redis_once는 한 번만 실행되도록 보장하는 sync.Once 인스턴스
	auth_redis_once.Do(func() {
		if auth_redis == nil {
			auth_redis = InitAuthRedis()
		}
	})
	return auth_redis
}
