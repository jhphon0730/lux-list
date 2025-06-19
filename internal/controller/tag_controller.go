package controller

import (
	"lux-list/internal/service"

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
