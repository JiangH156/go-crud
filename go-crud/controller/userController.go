package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"john/gin-curd/common"
	"john/gin-curd/dto"
	"john/gin-curd/models"
	"john/gin-curd/response"
	"john/gin-curd/util"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	user := models.User{}
	c.ShouldBind(&user)
	tx := db.Begin()

	// 当name不为空时，需要判断长度
	if user.Name != "" {
		length := util.GetStringLength(user.Name)
		if length >= 30 {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "账号名长度必须小于30")
			return
		}
	}

	if len(user.Password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于等于6位")
		return
	}
	if len(user.Telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话只能为11位")
		return
	}
	if len(user.Name) == 0 {
		user.Name = util.RandString(10)
	}

	////用户名已存在
	//if hasUser(db, user.Name) {
	//	response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
	//	return
	//}

	//已经存在该用户
	if isTelephone(db, user.Telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在或电话号码重复注册")
		return
	}

	fmt.Println(user.ID)

	//加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密失败")
		return
	}

	newUser := models.User{
		Name:      user.Name,
		Password:  string(hasedPassword),
		Telephone: user.Telephone,
	}
	// 添加用户
	tx.Create(&newUser)

	//生成token
	token, err := common.GenToken(newUser.ID)
	fmt.Println(token)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "生成token失败")
		log.Printf("token generate error: %v", err)
		tx.Rollback()
		return
	}

	response.Response(c, http.StatusOK, 200, gin.H{
		"user":  dto.ToUserDto(user),
		"token": token,
	}, "用户添加成功")
	tx.Commit()
}

func Login(ctx *gin.Context) {
	var DB = common.GetDB()
	var user models.User
	ctx.ShouldBind(&user)

	//获取参数
	telephone := user.Telephone
	password := user.Password
	//数据验证
	fmt.Println(telephone, "手机号码长度", len(telephone))
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var newUser models.User
	DB.Where("telephone = ?", telephone).First(&newUser)
	if newUser.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.GenToken(user.ID)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")

}

func isTelephone(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

func hasUser(db *gorm.DB, name string) bool {
	var user models.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}
