package dao

import (
	"context"
	"gorm.io/gorm"
	"sends/models"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}
func (dao *UserDao) ShowAllWait(Organization uint, Posts uint) (waits []*models.Wait, err error) {
	err = dao.DB.Model(&models.Wait{}).
		Where("Organization=? AND Posts=? AND State=?", Organization, Posts, 0).
		Find(&waits).
		Error
	return
}
func (dao *UserDao) UpdateState(Organization uint, Posts uint, stuNum string, State int) (err error) {
	err = dao.DB.Model(&models.Wait{}).
		Where("Organization=? AND Posts=? AND stu_num=?", Organization, Posts, stuNum).
		Update("State", State).
		Error
	return
}
