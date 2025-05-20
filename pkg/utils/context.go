// Controller의 Context에서 gin.Context를 사용하기 위한 패키지
package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Context에서 userID를 integer로 변환하여 가져오는 함수
func GetUserIDFromContext(ctx *gin.Context) (int, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return 0, errors.New("userID not found in context")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return 0, errors.New("userID is not of type int")
	}

	return userIDInt, nil
}
