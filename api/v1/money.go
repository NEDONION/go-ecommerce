package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
	"net/http"
)

func ShowMoney(c *gin.Context) {
	showMoneyService := service.ShowMoneyService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err == nil {
		res := showMoneyService.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, http.StatusBadRequest)
		util.LogrusObj.Infoln(err)
	}
}
