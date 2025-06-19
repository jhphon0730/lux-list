package service

import (
	"database/sql"
	"errors"
	"lux-list/internal/model"
	"lux-list/internal/repository"
	"net/http"
)

// TagService는 태그 관련 메서드를 정의하는 인터페이스
type TagService interface {
	GetTagsByTagID(userID int, tagID int) (*model.Tag, int, error)
	GetTagsByUserID(userID int) ([]model.Tag, int, error)
	GetTagsByTaskID(userID int, taskID int) ([]model.Tag, int, error)
}

// tagService는 TagService 인터페이스를 구현하는 구조체
type tagService struct {
	tagRepository repository.TagRepository
}

// NewTagService는 TagService의 인스턴스를 생성하는 함수
func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepository,
	}
}

// GetTagsByTagID는 사용자의 특정 태그를 조회하는 메서드
func (s *tagService) GetTagsByTagID(userID int, tagID int) (*model.Tag, int, error) {
	tag, err := s.tagRepository.GetTagsByTagID(userID, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("tag not found")
		}
		return nil, http.StatusInternalServerError, err
	}
	return tag, http.StatusOK, nil
}

// GetTagsByUserID는 사용자의 모든 태그를 조회하는 메서드
func (s *tagService) GetTagsByUserID(userID int) ([]model.Tag, int, error) {
	tags, err := s.tagRepository.GetTagsByUserID(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return tags, http.StatusOK, nil
}

// GetTagsByTaskID는 특정 작업에 연결된 태그를 조회하는 메서드
func (s *tagService) GetTagsByTaskID(userID int, taskID int) ([]model.Tag, int, error) {
	tags, err := s.tagRepository.GetTagsByTaskID(userID, taskID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return tags, http.StatusOK, nil
}
