package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sends/dao"
	"sends/models"
	"sends/utils"
)

// ShowAll
// @Tags 管理员
// @Summary 通过审核
// @Param token header string true "token"
// @Param messageId body models.AdminChooseTask true "获取信息id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/show [get]
func ShowAll(c *gin.Context) {
	auth := c.GetHeader("token")
	_, err := utils.AnalyseAdminToken(auth)
	if err != nil {
		log.Println("ShowAll Error", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ShowAll" + err.Error(),
			"data": "",
		})
		return
	}
	json := models.AdminChooseTask{}
	organization := json.Organization
	posts := json.Posts
	result := dao.NewUserDao(c.Request.Context())
	data, _ := result.ShowAllWait(organization, posts)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
		"data": data,
	})
}
