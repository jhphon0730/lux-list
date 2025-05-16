package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 서버 정보를 구성하는 구조체
type Server struct {
	Port string
	Ctx  context.Context

	Engine *gin.Engine
	Server *http.Server
}

// 서버 구조체 생성자 함수
func NewServer(Port string, ctx context.Context) *Server {
	// gin engine 초기화
	engine := gin.Default()
	engine.Use(gin.Logger())

	// http server 초기화
	server := &http.Server{
		Addr:    ":" + Port,
		Handler: engine,
	}

	return &Server{
		Port: Port,
		Ctx:  ctx,

		Engine: engine,
		Server: server,
	}
}
