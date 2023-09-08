package models

import (
	"gorm.io/gorm"
	"time"
)

//等待面试列表

type Wait struct {
	gorm.Model
	Organization     uint      //组织id
	OrganizationName string    //组织名
	Posts            uint      //职位id
	PostsName        string    //职位名
	Name             string    //姓名
	StuNum           string    //学号
	Instructions     string    //特别说明
	Time             time.Time //面试时间
	Address          string    //面试地点
	Path             string    //组织标识图片地址
	State            int       //状态 0：待面试（默认） 1：面试结束
}
type AdminInfo struct {
	Name     string `json:"name"`     //name
	Password string `json:"password"` //password
}
type AdminChooseTask struct {
	Organization uint   `json:"Organization"` //organization
	Posts        uint   `json:"Posts"`        //posts
	StuNum       string `json:"StuNum"`       //StuNum
	State        int    `json:"State"`        //State
}
