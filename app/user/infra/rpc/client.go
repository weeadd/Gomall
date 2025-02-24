package rpc

import (
	"Gomall/app/user/conf"
	"Gomall/app/user/utils"
	"Gomall/common/clientsuite"
	"Gomall/rpc_gen/kitex_gen/auth/authservice"
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	AuthClient   authservice.Client
	once         sync.Once
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
)

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

	AuthClient, err = authservice.NewClient("auth", opts...)
	utils.MustHandleError(err)
}
