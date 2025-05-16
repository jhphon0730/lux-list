package controller

import (
	"net/http"

	"lux-list/internal/middleware"
	"lux-list/internal/model"
	"lux-list/internal/service"

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

func RegisterRoutes(router *gin.RouterGroup, authController AuthController) {
	router.POST("/login", authController.Login)
	router.GET("/logout", authController.Logout)
	router.GET("/profile", middleware.AuthMiddleware(), authController.Profile)
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
	session.Set("user", user.ID)
	session.Set("access_token", token)

	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	ctx.JSON(status, gin.H{"message": "User Logged In", "user": user})
}

// logout은 사용자 로그아웃 요청을 처리하는 메서드
func (c *authController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "User Logged Out"})
}

// profile은 요청 사용자의 프로필 정보를 반환하는 메서드
func (c *authController) Profile(ctx *gin.Context) {
	userID, _ := ctx.Get("user")

	user, err := c.authService.GetUserByID(userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
