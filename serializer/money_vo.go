package serializer

import (
	"go-ecommerce/model"
	util "go-ecommerce/pkg/utils"
)

type MoneyVO struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) MoneyVO {
	util.Encrypt.SetKey(key)
	return MoneyVO{
		UserID:    item.ID,
		UserName:  item.UserName,
		UserMoney: util.Encrypt.AesDecoding(item.Money),
	}
}
