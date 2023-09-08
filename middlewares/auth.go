package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sends/config"
	"sends/utils"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("token")
		resp, err := utils.AnalyseAdminToken(auth)
		if err != nil {
			c.Abort()
			log.Printf("%+v", err)
			c.JSON(http.StatusOK,
				gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "Unauthorized",
				})
			return
		}
		if resp.Name != config.Name || resp.PassWord != config.Password {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "你不是管理员",
				"data": "",
			})
			return
		}
		c.Next()
	}
}
