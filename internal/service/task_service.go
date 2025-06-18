package service

import (
	"database/sql"
	"errors"
	"net/http"

	"lux-list/internal/model"
	"lux-list/internal/repository"
)

// TaskService는 작업 관련 메서드를 정의하는 인터페이스
type TaskService interface {
	GetTasks(userID int) ([]model.Task, int, error)
	GetTasksByTaskID(userID int, taskID int) (*model.Task, int, error)
	CreateTasks(userID int, task *model.Task) (*model.Task, int, error)
	DeleteTasks(userID int, taskID int) (int, error)
	UpdateTasks(userID int, taskID int, task *model.Task) (*model.Task, int, error)
	CompleteTasks(userID int, taskID int) (*model.Task, int, error)
	InCompleteTasks(userID int, taskID int) (*model.Task, int, error)
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

// GetTasksByTaskID는 사용자의 특정 작업을 조회하는 메서드
func (s *taskService) GetTasksByTaskID(userID int, taskID int) (*model.Task, int, error) {
	task, err := s.taskRepository.GetTasksByTaskID(userID, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("task not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	return task, http.StatusOK, nil
}

// CreateTasks는 사용자의 작업을 생성하는 매서드
func (s *taskService) CreateTasks(userID int, task *model.Task) (*model.Task, int, error) {
	created_task, err := s.taskRepository.CreateTasks(userID, task)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return created_task, http.StatusCreated, nil
}

// DeleteTasks는 사용자의 작업을 삭제하는 메서드
func (s *taskService) DeleteTasks(userID int, taskID int) (int, error) {
	err := s.taskRepository.DeleteTasks(userID, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("task not found")
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

// UpdateTasks는 사용자의 작업을 업데이트하는 메서드
func (s *taskService) UpdateTasks(userID int, taskID int, task *model.Task) (*model.Task, int, error) {
	updatedTask, err := s.taskRepository.UpdateTasks(userID, taskID, task)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return updatedTask, http.StatusOK, nil
}

// CompleteTasks는 사용자의 작업을 완료 상태로 변경하는 메서드
func (s *taskService) CompleteTasks(userID int, taskID int) (*model.Task, int, error) {
	task, err := s.taskRepository.GetTasksByTaskID(userID, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("task not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	task.IsCompleted = true
	updatedTask, err := s.taskRepository.UpdateTasks(userID, taskID, task)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return updatedTask, http.StatusOK, nil
}

// InCompleteTasks는 사용자의 작업을 완료 상태에서 미완료 상태로 변경하는 메서드
func (s *taskService) InCompleteTasks(userID int, taskID int) (*model.Task, int, error) {
	task, err := s.taskRepository.GetTasksByTaskID(userID, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("task not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	task.IsCompleted = false
	updatedTask, err := s.taskRepository.UpdateTasks(userID, taskID, task)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return updatedTask, http.StatusOK, nil
}
