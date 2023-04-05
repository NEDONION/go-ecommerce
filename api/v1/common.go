package v1

import (
	"encoding/json"
	"fmt"
	"go-ecommerce/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
