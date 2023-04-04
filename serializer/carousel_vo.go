package serializer

import "go-ecommerce/model"

type CarouselVO struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item *model.Carousel) CarouselVO {
	return CarouselVO{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化轮播图列表
func BuildCarousels(items []*model.Carousel) (carousels []CarouselVO) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
