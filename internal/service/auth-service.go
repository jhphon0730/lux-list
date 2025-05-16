package service

import (
	"net/http"

	"lux-list/internal/model"
	"lux-list/internal/repository"
	"lux-list/pkg/auth"
)

// AuthService는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthService interface {
	ExistUser(name string) (bool, error)
	Login(name string) (*model.User, string, int, error)
	RegisterAndGenerateJWT(name string) (*model.User, string, int, error)
	GetUserByName(name string) (*model.User, error)
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

// ExistUser는 사용자가 존재하는지 확인하는 메서드
func (s *authService) ExistUser(name string) (bool, error) {
	exist, err := s.authRepository.ExistUser(name)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// Login은 JWT 토큰 발급을 위해 사용자 로그인 요청을 처리하는 메서드
func (s *authService) Login(name string) (*model.User, string, int, error) {
	user, err := s.authRepository.GetUserByName(name)
	if err != nil {
		return nil, "", http.StatusBadRequest, err
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", http.StatusInternalServerError, err
	}

	return user, token, http.StatusOK, nil
}

// RegisterAndGenerateJWT는 새로운 사용자를 생성하고 JWT 토큰을 발급하는 메서드
func (s *authService) RegisterAndGenerateJWT(name string) (*model.User, string, int, error) {
	user, err := s.authRepository.CreateUser(name)
	if err != nil {
		return nil, "", http.StatusBadRequest, err
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", http.StatusInternalServerError, err
	}
	return user, token, http.StatusOK, nil
}

// GetUserByName은 사용자 이름으로 사용자를 조회하는 메서드
func (s *authService) GetUserByName(name string) (*model.User, error) {
	user, err := s.authRepository.GetUserByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
