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
		v1.GET("products", api.ListProducts)
		v1.GET("product/:id", api.ShowProductDetails)
		v1.POST("products", api.SearchProducts)
		v1.GET("product/images/:id", api.ListProductImg) // 获取商品图片
		v1.GET("categories", api.ListCategories)         //商品分类

		v1.GET("carousels", api.ListCarousels) // 轮播图

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar) // 上传头像

			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)

			// 收藏夹
			authed.GET("favorites", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

		}
	}
	return r
}
