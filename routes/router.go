package routes

import (
	"github.com/gin-gonic/gin"
	api "go-ecommerce/api/v1"
	"go-ecommerce/middleware"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 商品操作
		v1.GET("carousels", api.ListCarousels) //轮播图

		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar) //上传头像
		}
	}
	return r
}
