package middleware

import (
	"lux-list/pkg/auth"
	"lux-list/pkg/redis"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware는 인증 미들웨어를 정의하는 함수
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("user")
		accessToken := session.Get("access_token")
		if userID == nil || accessToken == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인이 필요한 서비스입니다."})
			ctx.Abort()
			return
		}

		// Redis 세션에 저장 된 access_token과 비교하여 검증 ( 중복 로그인 방지 )
		if auth_key, _ := redis.GetAuthSession(ctx, userID); auth_key == "" || auth_key != accessToken {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 세션이 만료되었습니다."})
			ctx.Abort()
			return
		}

		// JWT 검증 로직 추가
		claims, err := auth.ValidateAndParseJWT(accessToken.(string))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 세션이 만료되었습니다."})
			ctx.Abort()
			return
		}

		ctx.Set("user", claims.UserID)
		ctx.Set("access_token", accessToken)
		ctx.Next()
	}
}
