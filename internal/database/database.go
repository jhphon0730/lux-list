package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"lux-list/internal/config"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db_instance *sql.DB
	once        sync.Once
)

// 데이터베이스 초기화 함수, 데이터베이스 연결 및 마이그레이션 수행
func InitDB() error {
	config := config.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.Database.DB_HOST,
		config.Database.DB_PORT,
		config.Database.DB_USER,
		config.Database.DB_PASSWORD,
		config.Database.DB_NAME,
		config.Database.SSL_MODE,
		config.Database.TIMEZONE,
	)

	var err error
	db_instance, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}

	if err = db_instance.Ping(); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	// 마이그레이션 sql 파일 경로 찾기
	schemaPath := filepath.Join("internal", "database", "schema.sql")
	schema, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	// 마이그레이션 sql 실행
	stmts := strings.Split(string(schema), ";")
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db_instance.Exec(stmt)
		if err != nil {
			// "already exists" 에러는 무시
			if strings.Contains(err.Error(), "already exists") {
				continue
			}
			log.Printf("Migration error: %v\nSQL: %s\n", err, stmt)
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}

// Returns the global DB connection
func GetDB() *sql.DB {
	once.Do(func() {
		InitDB()
		if db_instance == nil {
			log.Fatal("DB instance is not initialized")
		}
	})
	return db_instance
}
