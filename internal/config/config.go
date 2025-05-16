package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// 서버의 정보를 구성하는 구조체
type ServerConfig struct {
	Host string
	Port string
}

// 프로그램의 환경변수 설정을 포함하는 구조체
type Config struct {
	Server ServerConfig
}

// Config 구조체의 인스턴스를 저장하기 위한 변수와 동기화 객체
var (
	config_instance *Config
	once            sync.Once
)

// .env 파일을 로드하여 구성 정보를 반환하는 함수
func loadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "5000"),
		},
	}
}

// Config 구조체의 인스턴스를 반환하는 함수
func GetConfig() *Config {
	once.Do(func() {
		config_instance = loadConfig()
		if config_instance == nil {
			log.Fatalln("Failed to load config (.env)")
		}
	})

	return config_instance
}

// 환경변수 값을 가져오는 함수
func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
