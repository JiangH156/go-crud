package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"john/gin-curd/common"
	"john/gin-curd/models"
	"john/gin-curd/util"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	user := models.User{}
	c.ShouldBind(&user)

	if len(user.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码必须大于等于6位",
		})
		return
	}
	if len(user.Telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "电话只能为11位",
		})
		return
	}
	if len(user.Name) == 0 {
		user.Name = util.RandString(10)
	}

	//已经存在该用户
	if isTelephone(db, user.Telephone) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "用户已存在或电话号码重复注册",
		})
		return
	}
	if db.Create(&user).Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":       "用户添加成功",
			"name":      user.Name,
			"password":  user.Password,
			"telephnoe": user.Telephone,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加用户失败",
		})
	}
}

func isTelephone(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
