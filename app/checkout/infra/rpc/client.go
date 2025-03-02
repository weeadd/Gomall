package rpc

import (
	"Gomall/app/checkout/conf"
	"Gomall/common/clientsuite"
	"Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"Gomall/rpc_gen/kitex_gen/order/orderservice"
	"Gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	OrderClient   orderservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func MustHandleError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}

func Init() {
	once.Do(func() {
		initClient()
	})
}

func initClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	OrderClient, err = orderservice.NewClient("order", opts...)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)

	if err != nil {
		hlog.Fatal(err)
	}

}
