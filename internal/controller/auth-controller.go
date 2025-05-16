package controller

import (
	"lux-list/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthController는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthController interface {
	login(c *gin.Context)
	logout(c *gin.Context)
}

// authController는 AuthController 인터페이스를 구현하는 구조체
type authController struct {
	authService service.AuthService
}

func RegisterRoutes(router *gin.RouterGroup, authController AuthController) {
	router.POST("/login", authController.login)
	router.GET("/logout", authController.logout)
}

// NewAuthController는 AuthController의 인스턴스를 생성하는 함수
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// login은 사용자 로그인 요청을 처리하는 메서드
func (c *authController) login(ctx *gin.Context) {}

// logout은 사용자 로그아웃 요청을 처리하는 메서드
func (c *authController) logout(ctx *gin.Context) {}
