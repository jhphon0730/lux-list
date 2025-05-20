package repository

import "lux-list/internal/model"

type TaskRepository interface {
	GetTasks(userID int) ([]model.Task, error)
	GetTaskByID(taskID int) (*model.Task, error)
	CreateTask(userID int, task *model.Task) (*model.Task, error)
}

type taskRepository struct{}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

// GetTasks는 사용자의 모든 작업을 조회하는 메서드
func (r *taskRepository) GetTasks(userID int) ([]model.Task, error) {
	var tasks []model.Task
	return tasks, nil
}

// GetTaskByID는 작업 ID로 작업을 조회하는 메서드
func (r *taskRepository) GetTaskByID(taskID int) (*model.Task, error) {
	var task model.Task
	return &task, nil
}

// CreateTask는 새로운 작업을 생성하는 메서드
func (r *taskRepository) CreateTask(userID int, task *model.Task) (*model.Task, error) {
	return task, nil
}
