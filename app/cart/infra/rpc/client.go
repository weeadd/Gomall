package rpc

import (
	"Gomall/app/cart/conf"
	"Gomall/app/cart/utils"
	"Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient cartservice.Client
	ProductClient productcatalogservice.Client
	once sync.Once
) 

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])//"127.0.0.1:8500"
	utils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("product", client.WithResolver(r))
	utils.MustHandleError(err)
}