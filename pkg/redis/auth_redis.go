package redis

import (
	"context"
	"errors"
	"sync"
	"time"

	"lux-list/internal/config"
	"lux-list/pkg/utils"

	"github.com/go-redis/redis/v8"
)

type AuthRedisClient = redis.Client

const (
	authSessionKey = "auth_session:"
)

var (
	// auth_redis는 Redis 클라이언트 인스턴스
	auth_redis *AuthRedisClient
	// auth_redis_once는 Redis 클라이언트 초기화를 위한 sync.Once 인스턴스
	auth_redis_once sync.Once
)

// InitAuthRedis는 Redis 클라이언트를 초기화하고 반환하는 함수
func InitAuthRedis(ctx context.Context) (*AuthRedisClient, error) {
	config := config.GetConfig()
	auth_redis = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       utils.InterfaceToInt(config.Redis.AuthDB),
	})

	// Ping 메서드를 사용하여 Redis 서버에 연결을 확인
	_, err := auth_redis.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return auth_redis, nil
}

// GetAuthRedis는 Redis 클라이언트를 반환하는 함수
func GetAuthRedis(ctx context.Context) (*AuthRedisClient, error) {
	// auth_redis_once는 한 번만 실행되도록 보장하는 sync.Once 인스턴스
	auth_redis_once.Do(func() {
		if auth_redis == nil {
			auth_redis, _ = InitAuthRedis(ctx)
		}
	})
	if auth_redis == nil {
		return nil, errors.New("redis client is not initialized")
	}
	return auth_redis, nil
}

// SetAuthSession은 Redis에 인증 세션을 저장하는 함수
func SetAuthSession(ctx context.Context, key interface{}, value interface{}) error {
	set_key := authSessionKey + utils.InterfaceToString(key)

	// TTL 1시간
	err := auth_redis.Set(ctx, set_key, value, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

// DeleteAuthSession은 Redis에서 인증 세션을 삭제하는 함수
func DeleteAuthSession(ctx context.Context, key interface{}) error {
	set_key := authSessionKey + utils.InterfaceToString(key)

	err := auth_redis.Del(ctx, set_key).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetAuthSession은 Redis에서 인증 세션을 가져오는 함수
func GetAuthSession(ctx context.Context, key interface{}) (string, error) {
	set_key := authSessionKey + utils.InterfaceToString(key)

	val, err := auth_redis.Get(ctx, set_key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
