package rpc

import (
	"Gomall/app/cart/conf"
	"Gomall/app/cart/utils"
	"Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"Gomall/common/clientsuite"
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	CartClient 		cartservice.Client
	ProductClient 	productcatalogservice.Client
	once 			sync.Once
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err				error
) 

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	 opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	 }

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)
}