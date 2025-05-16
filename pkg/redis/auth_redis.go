package redis

// AuthRedis는 인증 관련 Redis 작업을 정의하는 인터페이스
type AuthRedis interface{}

// authRedis는 Redis를 사용한 인증 관련 작업을 구현하는 구조체
type authRedis struct{}

func NewAuthRedis() AuthRedis {
	return &authRedis{}
}
