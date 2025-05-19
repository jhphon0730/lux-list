package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"lux-list/internal/config"
	"lux-list/internal/database"
	"lux-list/internal/server"
	"lux-list/pkg/redis"
)

func main() {
	// 로그 설정 및 컨텍스트 생성
	ctx, cancel := context.WithCancel(context.Background())
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// config 파일 로딩 (.env)
	config := config.GetConfig()
	if config == nil {
		log.Fatal("Failed to load config")
	}

	// 데이터베이스 초기화
	if err := database.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Redis 초기화
	if err := redis.InitRedis(ctx); err != nil {
		log.Fatalf("Redis initialization failed: %v", err)
	}

	// 서버 생성
	srv := server.NewServer(config.Server.Port, ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// (고루틴) 서버 실행
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 프로그램 종료 대기
	<-c
	log.Println("Shutdown signal received")

	// 서버 종료
	srv.Shutdown()
	cancel()
	log.Println("Server shutdown complete")
}
