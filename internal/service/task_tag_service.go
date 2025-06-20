package service

import (
	"database/sql"
	"errors"
	"lux-list/internal/model"
	"lux-list/internal/repository"
	"net/http"
)

// TaskTagService는 작업 태그 관련 메서드를 정의하는 인터페이스
type TaskTagService interface {
	AddTagToTask(taskID int, tagID int) (int, error)
	RemoveTagFromTask(taskID int, tagID int) (int, error)
	GetTagsByTaskID(taskID int) ([]model.Tag, int, error)
}

// taskTagService는 TaskTagService 인터페이스를 구현하는 구조체
type taskTagService struct {
	taskTagRepository repository.TaskTagRepository
}

// NewTaskTagService는 TaskTagService의 인스턴스를 생성하는 함수
func NewTaskTagService(taskTagRepository repository.TaskTagRepository) TaskTagService {
	return &taskTagService{
		taskTagRepository: taskTagRepository,
	}
}

// AddTagToTask는 작업에 태그를 추가하는 메서드
func (s *taskTagService) AddTagToTask(taskID int, tagID int) (int, error) {
	err := s.taskTagRepository.AddTagToTask(taskID, tagID)
	if err != nil {
		if err == sql.ErrTxDone {
			return http.StatusConflict, errors.New("tag already exists in task")
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// RemoveTagFromTask는 작업에서 태그를 제거하는 메서드
func (s *taskTagService) RemoveTagFromTask(taskID int, tagID int) (int, error) {
	err := s.taskTagRepository.RemoveTagFromTask(taskID, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("tag not found in task")
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// GetTagsByTaskID는 특정 작업에 연결된 태그를 조회하는 메서드
func (s *taskTagService) GetTagsByTaskID(taskID int) ([]model.Tag, int, error) {
	tags, err := s.taskTagRepository.GetTagsByTaskID(taskID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return tags, http.StatusOK, nil
}
