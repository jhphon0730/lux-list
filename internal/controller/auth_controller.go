package controller

import (
	"fmt"
	"net/http"

	"lux-list/internal/middleware"
	"lux-list/internal/model"
	"lux-list/internal/service"
	"lux-list/pkg/redis"
	"lux-list/pkg/types"
	"lux-list/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthController는 사용자 인증 관련 메서드를 정의하는 인터페이스
type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Profile(c *gin.Context)
}

// authController는 AuthController 인터페이스를 구현하는 구조체
type authController struct {
	authService service.AuthService
}

// RegisterRoutes는 인증 관련 라우트를 등록하는 함수
func RegisterAuthRoutes(router *gin.RouterGroup, authController AuthController) {
	router.POST("/login", authController.Login)
	router.GET("/logout", middleware.AuthMiddleware(), authController.Logout)
	router.GET("", middleware.AuthMiddleware(), authController.Profile)
}

// NewAuthController는 AuthController의 인스턴스를 생성하는 함수
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// login은 사용자 로그인 요청을 처리하는 메서드
func (c *authController) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 사용자 이름이 비어있는지 확인
	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	isExistUser, err := c.authService.ExistUser(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(ctx)
	var (
		user   *model.User
		token  string
		status int
	)

	if isExistUser {
		user, token, status, err = c.authService.Login(req.Name)
	} else {
		user, token, status, err = c.authService.RegisterAndGenerateJWT(req.Name)
	}
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// 세션에 사용자 정보와 토큰 저장
	session.Set(types.SESSION_USERID, user.ID)
	session.Set(types.SESSION_ACCESS_TOKEN, token)

	// Redis에 세션 저장
	if err := redis.SetAuthSession(ctx, user.ID, token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session in Redis"})
		return
	}

	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	ctx.JSON(status, gin.H{"message": "User Logged In", "user": user})
}

// logout은 사용자 로그아웃 요청을 처리하는 메서드
func (c *authController) Logout(ctx *gin.Context) {
	userID, _ := utils.GetUserIDFromContext(ctx)
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()

	fmt.Print(userID)
	if err := redis.DeleteAuthSession(ctx, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session in Redis"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Logged Out"})
}

// profile은 요청 사용자의 프로필 정보를 반환하는 메서드
func (c *authController) Profile(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, status, err := c.authService.GetUserByID(userID)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
