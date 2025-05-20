package service

import (
	"lux-list/internal/model"
	"lux-list/internal/repository"
	"net/http"
)

// TaskService는 작업 관련 메서드를 정의하는 인터페이스
type TaskService interface {
	GetTasks(userID int) ([]model.Task, int, error)
}

// taskService는 TaskService 인터페이스를 구현하는 구조체
type taskService struct {
	taskRepository repository.TaskRepository
}

// NewTaskService는 TaskService의 인스턴스를 생성하는 함수
func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

// GetTasks는 사용자의 모든 작업을 조회하는 메서드
func (s *taskService) GetTasks(userID int) ([]model.Task, int, error) {
	tasks, err := s.taskRepository.GetTasks(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return tasks, http.StatusOK, nil
}
