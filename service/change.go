package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sends/dao"
	"sends/models"
	"sends/utils"
)

// ChangState
// @Tags 管理员
// @Summary 通过审核
// @Param token header string true "token"
// @Param messageId body models.AdminChooseTask true "获取信息id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/change [post]

func ChangeState(c *gin.Context) {
	auth := c.GetHeader("token")
	_, err := utils.AnalyseAdminToken(auth)
	if err != nil {
		log.Println("ChangeState Error", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ChangeState" + err.Error(),
			"data": "",
		})
		return
	}
	json := models.AdminChooseTask{}
	organization := json.Organization
	posts := json.Posts
	stuNum := json.StuNum
	state := json.State
	result := dao.NewUserDao(c.Request.Context())
	data := result.UpdateState(organization, posts, stuNum, state)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
		"data": data,
	})
}
