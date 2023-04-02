package dao

import (
	"fmt"
	"go-ecommerce/model"
	"os"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{},
			&model.Product{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.ProductImg{},
			&model.Order{},
			&model.Cart{},
			&model.Admin{},
			&model.Address{},
			&model.Notice{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	return
}
