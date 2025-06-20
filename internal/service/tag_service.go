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
	CreateTags(userID int, tag *model.Tag) (*model.Tag, int, error)
	DeleteTags(userID int, tagID int) (int, error)
	UpdateTags(userID int, tagID int, tag *model.Tag) (*model.Tag, int, error)
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

// GetTagsByTagID는 태그 ID로 태그를 조회하는 메서드
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

// CreateTags는 사용자의 태그를 생성하는 메서드
func (s *tagService) CreateTags(userID int, tag *model.Tag) (*model.Tag, int, error) {
	createdTag, err := s.tagRepository.CreateTags(userID, tag)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return createdTag, http.StatusCreated, nil
}

// DeleteTags는 태그를 삭제하는 메서드
func (s *tagService) DeleteTags(userID int, tagID int) (int, error) {
	err := s.tagRepository.DeleteTags(userID, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("tag not found")
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

// UpdateTags는 태그를 업데이트하는 메서드
func (s *tagService) UpdateTags(userID int, tagID int, tag *model.Tag) (*model.Tag, int, error) {
	updatedTag, err := s.tagRepository.UpdateTags(userID, tagID, tag)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return updatedTag, http.StatusOK, nil
}
