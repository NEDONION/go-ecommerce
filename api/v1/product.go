package v1

import (
	"github.com/gin-gonic/gin"
	util "go-ecommerce/pkg/utils"
	"go-ecommerce/service"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListProducts 展示商品列表
func ListProducts(c *gin.Context) {
	listProductsService := service.ProductService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.ListProducts(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// SearchProducts 搜索商品
func SearchProducts(c *gin.Context) {
	searchProductsService := service.ProductService{}
	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.SearchProducts(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowProductDetails 商品详情
func ShowProductDetails(c *gin.Context) {
	showProductService := service.ProductService{}
	res := showProductService.ShowDetails(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// ListProductImg 获取商品图片
func ListProductImg(c *gin.Context) {
	var listProductImgService service.ListProductImgService
	if err := c.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.List(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
