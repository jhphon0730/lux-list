package service

import "lux-list/internal/repository"

// AuthService는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthService interface {
}

// authService는 AuthService 인터페이스를 구현하는 구조체
type authService struct {
	authRepository repository.AuthRepository
}

// NewAuthService는 AuthService의 인스턴스를 생성하는 함수
func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}
