package controller

import (
	"lux-list/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthController는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthController interface{}

// authController는 AuthController 인터페이스를 구현하는 구조체
type authController struct {
	authService service.AuthService
}

func RegisterRoutes(router *gin.RouterGroup, authController AuthController) {
	router.POST("/login", nil)
	router.GET("/logout", nil)
}

// NewAuthController는 AuthController의 인스턴스를 생성하는 함수
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}
