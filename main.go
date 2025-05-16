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
	ctx, cancle := context.WithCancel(context.Background())
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server.NewServer("8080", ctx)
	// 고루틴 서버 실행
	go func() {
	}()

	// 프로그램 종료 대기
	<-c

	// HERE shutdown function
	cancle()
	log.Println("Shutdown signal received")
}
