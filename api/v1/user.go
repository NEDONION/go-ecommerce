package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
	"net/http"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, http.StatusBadRequest)
		util.LogrusObj.Infoln(err)
	}
}

// UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, http.StatusBadRequest)
		util.LogrusObj.Infoln(err)
	}
}

// UserUpdate 用户更新接口
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.Update(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, http.StatusBadRequest)
		util.LogrusObj.Infoln(err)
	}
}
