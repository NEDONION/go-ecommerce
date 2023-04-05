package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"go-ecommerce/dao"
	"go-ecommerce/model"
	"go-ecommerce/pkg/e"
	"go-ecommerce/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

// ProductService 商品服务
type ProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	model.BasePage
}

type ListProductImgService struct {
}

// Create 创建商品
func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadToQiNiu(tmp, files[0].Size)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryID),
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		Num:           service.Num,
		OnSale:        true,
		BossID:        int(uId),
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		productImgDao := dao.NewProductImgDao(ctx)
		tmp, _ = file.Open()
		path, err = UploadToQiNiu(tmp, file.Size)
		if err != nil {
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Data:   e.GetMsg(code),
				Error:  path,
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}

	wg.Wait()

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

// ShowDetails 展示商品详情
func (service *ProductService) ShowDetails(ctx context.Context, id string) serializer.Response {
	code := e.SUCCESS

	pId, _ := strconv.Atoi(id)

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

// ListProducts 展示商品列表
func (service *ProductService) ListProducts(ctx context.Context) serializer.Response {
	var products []*model.Product
	var total int64
	code := e.SUCCESS

	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["category_id"] = service.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
