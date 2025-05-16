package controller

import (
	"lux-list/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// authController는 AuthController 인터페이스를 구현하는 구조체
type authController struct {
	authService service.AuthService
}

func RegisterRoutes(router *gin.RouterGroup, authController AuthController) {
	router.POST("/login", authController.Login)
	router.GET("/logout", authController.Logout)
}

// NewAuthController는 AuthController의 인스턴스를 생성하는 함수
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// login은 사용자 로그인 요청을 처리하는 메서드
func (c *authController) Login(ctx *gin.Context) {
	type loginRequest struct {
		Name string `json:"name"`
	}

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	is_exist_user, err := c.authService.ExistUser(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !is_exist_user {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User exists"})
	}
}

// logout은 사용자 로그아웃 요청을 처리하는 메서드
func (c *authController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
