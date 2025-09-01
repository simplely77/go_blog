package middleware

import (
	"github.com/gin-gonic/gin"
	"server/model/appTypes"
	"server/model/response"
	"server/utils"
)

// AdminAuth 检查用户是否有管理员权限
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := utils.GetRoleID(c)
		if roleID != appTypes.Admin {
			response.Forbidden("Access denied. Admin privileges are required", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
