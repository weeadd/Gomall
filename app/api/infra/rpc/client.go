package rpc

import (
	"Gomall/app/api/conf"
	apiutils "Gomall/app/api/utils"
	"Gomall/rpc_gen/kitex_gen/auth/authservice"
	"Gomall/rpc_gen/kitex_gen/order/orderservice"
	"Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"Gomall/rpc_gen/kitex_gen/user/userservice"
	"Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	AuthClient    authservice.Client
	CartClient	  cartservice.Client
	once          sync.Once
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initOrderClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	apiutils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	apiutils.MustHandleError(err)

	AuthClient, err = authservice.NewClient("auth", client.WithResolver(r))
	apiutils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	apiutils.MustHandleError(err)

	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	apiutils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	apiutils.MustHandleError(err)

	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	apiutils.MustHandleError(err)
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	apiutils.MustHandleError(err)

	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	apiutils.MustHandleError(err)
}
