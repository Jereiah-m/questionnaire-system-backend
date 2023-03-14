package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire-system-backend/helper"
)

func AuthCheck() gin.HandlerFunc { //登录验证
	return func(c *gin.Context) {
		token := c.GetHeader("token")                 //获取token
		userClaims, err := helper.AnalyseToken(token) //解析token
		if err != nil {
			c.Abort() //终止后续处理
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "用户认证不通过",
			})
			return
		}
		c.Set("user_Claims", userClaims) //设置用户信息
		c.Next()                         //继续后续处理
	}
}
