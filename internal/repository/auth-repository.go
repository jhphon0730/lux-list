package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

type AuthRepository interface {
	ExistUser(name string) (bool, error)
	CreateUser(name string) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

// ExistUser는 사용자가 존재하는지 확인하는 메서드
func (r *authRepository) ExistUser(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE name = $1"
	row := r.db.QueryRow(query, name)

	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetUserByName은 사용자 이름으로 사용자를 조회하는 메서드
func (r *authRepository) GetUserByName(name string) (*model.User, error) {
	query := "SELECT id, name, created_at FROM users WHERE name = $1"
	row := r.db.QueryRow(query, name)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 사용자 없음
		}
		return nil, err // 다른 에러
	}

	return &user, nil
}

// CreateUser는 새로운 사용자를 생성하는 메서드
func (r *authRepository) CreateUser(name string) (*model.User, error) {
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id, name, created_at"
	row := r.db.QueryRow(query, name)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID는 사용자 ID로 사용자를 조회하는 메서드
func (r *authRepository) GetUserByID(id int) (*model.User, error) {
	query := "SELECT id, name, created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 사용자 없음
		}
		return nil, err // 다른 에러
	}

	return &user, nil
}
