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
}

// taskController는 TaskController 인터페이스를 구현하는 구조체
type taskController struct {
	taskService service.TaskService
}

// RegisterTaskRoutes는 작업 관련 라우트를 등록하는 함수
func RegisterTaskRoutes(router *gin.RouterGroup, taskController TaskController) {
	router.GET("", taskController.GetTasks)
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

	tasks, status, err := c.taskService.GetTasks(userID)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"tasks": tasks})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 입력 값이 유효한지 검사
	if err := req.CheckValidCreateTaskRequest(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ToTask(userID)
}
