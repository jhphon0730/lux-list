package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"lux-list/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // 오타 수정
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 서버 생성
	srv := server.NewServer("5000", ctx)

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
