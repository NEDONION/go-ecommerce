package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
	"net/http"
)

func ListCarousels(c *gin.Context) {
	listCarouselsService := service.ListCarouselsService{}
	if err := c.ShouldBind(&listCarouselsService); err == nil {
		res := listCarouselsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, http.StatusBadRequest)
		util.LogrusObj.Infoln(err)
	}
}
