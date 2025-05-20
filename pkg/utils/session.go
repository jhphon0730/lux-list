package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 세션을 지워주는 함수
func ClearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()
}
