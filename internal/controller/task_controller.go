package controller

import (
	"lux-list/internal/model"
	"lux-list/internal/service"
	"lux-list/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TaskController는 작업 관련 메서드를 정의하는 인터페이스
type TaskController interface {
	GetTasks(c *gin.Context)
	GetTasksByTaskID(c *gin.Context)
	CreateTasks(c *gin.Context)
	DeleteTasks(c *gin.Context)
	UpdateTasks(c *gin.Context)
	CompleteTasks(c *gin.Context)
	InCompleteTasks(c *gin.Context)
}

// taskController는 TaskController 인터페이스를 구현하는 구조체
type taskController struct {
	taskService service.TaskService
}

// RegisterTaskRoutes는 작업 관련 라우트를 등록하는 함수
func RegisterTaskRoutes(router *gin.RouterGroup, taskController TaskController) {
	router.GET("", taskController.GetTasks)
	router.GET("/:taskID", taskController.GetTasksByTaskID)
	router.POST("", taskController.CreateTasks)
	router.DELETE("/:taskID", taskController.DeleteTasks)
	router.PUT("/:taskID", taskController.UpdateTasks)
	router.PATCH("/:taskID/complete", taskController.CompleteTasks)
	router.PATCH("/:taskID/incomplete", taskController.InCompleteTasks)
}

// NewTaskController는 TaskController의 인스턴스를 생성하는 함수
func NewTaskController(taskService service.TaskService) TaskController {
	return &taskController{
		taskService: taskService,
	}
}

// GetTasks는 사용자의 모든 작업을 조회하는 메서드
func (c *taskController) GetTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	search_query := utils.GetTasksSearchQuery(ctx)
	taskListResult, status, err := c.taskService.GetTasks(userID, search_query)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"tasks": taskListResult.Tasks, "total_count": taskListResult.TotalCount})
}

// GetTasksByTaskID는 사용자의 특정 작업을 조회하는 메서드
func (c *taskController) GetTasksByTaskID(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskID := ctx.Param("taskID")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	task, status, err := c.taskService.GetTasksByTaskID(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"task": task})
}

// CreateTasks는 사용자의 작업을 생성하는 메서드
func (c *taskController) CreateTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req model.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing fields"})
		return
	}

	// 입력 값이 유효한지 검사
	if err := req.CheckValidCreateTaskRequest(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, status, err := c.taskService.CreateTasks(userID, req.ToTask(userID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"task": createdTask})
}

// DeleteTasks는 사용자의 작업을 삭제하는 메서드
func (c *taskController) DeleteTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskID := ctx.Param("taskID")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	status, err := c.taskService.DeleteTasks(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// UpdateTasks는 사용자의 작업을 업데이트하는 메서드
func (c *taskController) UpdateTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskID := ctx.Param("taskID")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	var req model.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing fields"})
		return
	}

	// 입력 값이 유효한지 검사
	if err := req.CheckValidUpdateTaskRequest(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findTask, status, err := c.taskService.GetTasksByTaskID(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	updatedTask, status, err := c.taskService.UpdateTasks(userID, utils.InterfaceToInt(taskID), req.ToTask(findTask))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"task": updatedTask})
}

// CompleteTasks는 사용자의 작업을 완료 상태로 업데이트하는 메서드
func (c *taskController) CompleteTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskID := ctx.Param("taskID")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	// 사용자의 작업을 완료 상태로 업데이트
	updatedTask, status, err := c.taskService.CompleteTasks(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"task": updatedTask})
}

// InCompleteTasks는 사용자의 작업을 미완료 상태로 업데이트하는 메서드
func (c *taskController) InCompleteTasks(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskID := ctx.Param("taskID")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	// 사용자의 작업을 미완료 상태로 업데이트
	updatedTask, status, err := c.taskService.InCompleteTasks(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"task": updatedTask})
}
