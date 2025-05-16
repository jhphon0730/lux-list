package server

import (
	"lux-list/internal/controller"
	"lux-list/internal/database"
	"lux-list/internal/repository"
	"lux-list/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	db = database.GetDB()

	authRepository = repository.NewAuthRepository(db)
	authService    = service.NewAuthService(authRepository)

	authController = controller.NewAuthController(authService)
)

// registerRoutes는 gin 엔진에 라우트를 등록하는 함수
func registerRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			controller.RegisterRoutes(auth, authController)
		}
	}
}
