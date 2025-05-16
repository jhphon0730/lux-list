package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userName := session.Get("user")
		accessToken := session.Get("access_token")
		if userName == nil || accessToken == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인이 필요한 서비스입니다."})
			ctx.Abort()
			return
		}

		ctx.Set("user", userName)
		ctx.Set("access_token", accessToken)
		ctx.Next()
	}
}
