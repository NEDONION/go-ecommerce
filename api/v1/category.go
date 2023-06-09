package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
)

func ListCategories(c *gin.Context) {
	listCategoriesService := service.ListCategoriesService{}
	if err := c.ShouldBind(&listCategoriesService); err == nil {
		res := listCategoriesService.ListCategories(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
