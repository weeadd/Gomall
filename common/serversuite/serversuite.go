package serversuite

import (
	"Gomall/common/mtl"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("",
			"",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry)),
		),
	}
	// consul
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}

	opts = append(opts, server.WithRegistry(r))
	return opts
}
