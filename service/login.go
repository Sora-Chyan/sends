package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sends/config"
	"sends/models"
	"sends/utils"
)

// AdminLogin
// @Tags 管理员
// @Summary 管理员登陆
// @Param adminInfo body models.AdminInfo true "管理员登陆"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/login [post]

func AdminLogin(c *gin.Context) {
	json := models.AdminInfo{}
	c.BindJSON(&json)
	name := json.Name
	password := json.Password
	log.Printf(name)
	log.Printf(password)
	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if name != config.Name || password != config.Password {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
			"data": "",
		})
		return
	}
	token := utils.GenerateAdminToken(name, password)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登陆成功",
		"data": gin.H{
			"token": token,
		},
	})

}
