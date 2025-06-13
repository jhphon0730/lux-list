package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"lux-list/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
	Shutdown() error
}

// 서버 정보를 구성하는 구조체
type server struct {
	Port string
	Ctx  context.Context

	Engine *gin.Engine
	Server *http.Server
}

// 서버 구조체 생성자 함수
func NewServer(Port string, ctx context.Context) Server {
	// gin engine 초기화
	engine := gin.Default()
	// engine.Use(gin.Logger())

	// http server 초기화
	httpSrv := &http.Server{
		Addr:    ":" + Port,
		Handler: engine,
	}

	return &server{
		Port: Port,
		Ctx:  ctx,

		Engine: engine,
		Server: httpSrv,
	}
}

// gin 엔진 설정 및 서버를 실행하는 함수
func (s *server) Run() error {
	// CORS 설정
	s.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.0.5:3000", "http://localhost:3000", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// sessions 설정
	session_store := cookie.NewStore([]byte(config.GetConfig().Server.SessionKey))
	session_store.Options(sessions.Options{
		MaxAge:   3600, // 1시간
		HttpOnly: true,
		Secure:   false, // 배포 시에는 true로 변경
		Path:     "/api/v1",
	})
	s.Engine.Use(sessions.Sessions("ss-token", session_store))

	// OPTIONS 설정
	s.Engine.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 라우트 등록
	registerRoutes(s.Engine)

	// 서버 실행
	if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// 서버를 종료하는 함수
func (s *server) Shutdown() error {
	log.Println("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(s.Ctx, 5*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
