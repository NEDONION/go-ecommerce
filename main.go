package main

import (
	"go-ecommerce/conf"
	"go-ecommerce/routes"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
