package rpc

import (
	"Gomall/app/api/conf"
	apiutils "Gomall/app/api/utils"
	"Gomall/rpc_gen/kitex_gen/user/userservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	apiutils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	apiutils.MustHandleError(err)
}
