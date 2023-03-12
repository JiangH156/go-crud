package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"john/gin-curd/common"
	"john/gin-curd/models"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取Authorization
		tokenString := c.GetHeader("Authorization")

		fmt.Print(tokenString)
		segs := strings.SplitN(tokenString, " ", 2)
		if len(segs) != 2 && segs[0] != "Beaber" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "授权参数有错",
			})
			c.Abort()
			return
		}

		// 判断token是否正确
		claims, err := common.ParseToken(segs[1])
		if err != nil {
			c.JSON(404, gin.H{
				"msg": err,
			})
			c.Abort()
			return
		}
		userId := claims.UserId
		var db = common.GetDB()
		var user models.User
		db.First(&user, userId)
		if user.ID == 0 {
			c.JSON(404, gin.H{
				"msg": "该用户不存在",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
