package middleware

import (
	"lux-list/pkg/auth"
	"lux-list/pkg/redis"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_USERID       = "userID"
	SESSION_ACCESS_TOKEN = "access_token"
)

// 세션을 지워주는 함수
func clearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()
}

// AuthMiddleware는 인증 미들웨어를 정의하는 함수
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get(SESSION_USERID)
		accessToken := session.Get(SESSION_ACCESS_TOKEN)
		if userID == nil || accessToken == nil {
			clearSession(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인이 필요한 서비스입니다."})
			ctx.Abort()
			return
		}

		accessTokenStr, ok := accessToken.(string)
		if !ok {
			clearSession(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "세션 토큰 형식 오류"})
			ctx.Abort()
			return
		}

		if authKey, err := redis.GetAuthSession(ctx, userID); err != nil || authKey == "" || authKey != accessTokenStr {
			clearSession(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 세션이 만료되었습니다."})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateAndParseJWT(accessTokenStr)
		if err != nil {
			clearSession(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 세션이 만료되었습니다."})
			ctx.Abort()
			return
		}

		ctx.Set(SESSION_USERID, claims.UserID)
		ctx.Set(SESSION_ACCESS_TOKEN, accessTokenStr)
		ctx.Next()
	}
}
