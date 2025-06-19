package controller

import (
	"lux-list/internal/service"
	"lux-list/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TagController는 태그 관련 메서드를 정의하는 인터페이스
type TagController interface {
	GetTagsByTagID(c *gin.Context)
	GetTagsByUserID(c *gin.Context)
	GetTagsByTaskID(c *gin.Context)
}

// tagController는 TagController 인터페이스를 구현하는 구조체
type tagController struct {
	tagService service.TagService
}

// RegisterTagRoutes는 태그 관련 라우트를 등록하는 함수
func RegisterTagRoutes(router *gin.RouterGroup, tagController TagController) {
	router.GET("/:tagID", tagController.GetTagsByTagID)
	router.GET("/user/:userID", tagController.GetTagsByUserID)
	router.GET("/task/:taskID", tagController.GetTagsByTaskID)
}

// NewTagController는 TagController의 인스턴스를 생성하는 함수
func NewTagController(tagService service.TagService) TagController {
	return &tagController{
		tagService: tagService,
	}
}

// GetTagsByTagID는 태그 ID로 태그를 조회하는 메서드
func (c *tagController) GetTagsByTagID(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tagID := ctx.Param("tagID")
	if tagID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tag ID is required"})
		return
	}

	tag, status, err := c.tagService.GetTagsByTagID(userID, utils.InterfaceToInt(tagID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(status, gin.H{"tag": tag})
}

// GetTagsByUserID는 사용자의 모든 태그를 조회하는 메서드
func (c *tagController) GetTagsByUserID(ctx *gin.Context) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tags, status, err := c.tagService.GetTagsByUserID(userID)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(status, gin.H{"tags": tags})
}

// GetTagsByTaskID는 작업 ID로 태그를 조회하는 메서드
func (c *tagController) GetTagsByTaskID(ctx *gin.Context) {
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

	tags, status, err := c.tagService.GetTagsByTaskID(userID, utils.InterfaceToInt(taskID))
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(status, gin.H{"tags": tags})
}
