package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
)

// CreateFavorite 创建收藏夹
func CreateFavorite(c *gin.Context) {
	favoritesService := service.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&favoritesService); err == nil {
		res := favoritesService.CreateFavorites(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowFavorites 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	favoritesService := service.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&favoritesService); err == nil {
		res := favoritesService.ShowFavorites(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteFavorite 删除收藏夹
func DeleteFavorite(c *gin.Context) {
	favoritesService := service.FavoritesService{}
	if err := c.ShouldBind(&favoritesService); err == nil {
		res := favoritesService.DeleteFavorites(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
