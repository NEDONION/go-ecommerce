package serializer

import "go-ecommerce/model"

type ProductImgVO struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImgVO {
	return ProductImgVO{
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImgVO) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
