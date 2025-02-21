package rpc

import (
	"Gomall/app/user/conf"
	"Gomall/app/user/utils"
	"Gomall/rpc_gen/kitex_gen/auth/authservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthClient authservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initClient()
	})
}

func initClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)

	AuthClient, err = authservice.NewClient("auth", client.WithResolver(r))
	utils.MustHandleError(err)
}
