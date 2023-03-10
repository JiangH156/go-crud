package controller

import (
	"github.com/gin-gonic/gin"
	"john/gin-curd/util"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")
	if len(password) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码不能为空",
		})
		return
	}
	if len(telephone) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "电话不能为空",
		})
		return
	}
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "电话只能为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码必须大于等于6位",
		})
		return
	}
	if len(username) == 0 {
		username = util.RandString(10)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": username,
	})
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
