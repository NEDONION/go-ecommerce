package service

import (
	"context"
	"go-ecommerce/dao"
	"go-ecommerce/serializer"
	"strconv"
)

// List 获取商品列表图片
func (service *ListProductImgService) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImgByProductId(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
