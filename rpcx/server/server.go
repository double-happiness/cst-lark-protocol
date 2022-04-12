package server

import (
	"cst-lark-protocal/utils"
	"github.com/rpcxio/rpcx-nacos/serverplugin"
	"github.com/smallnest/rpcx/server"
	"sync"
)

var (
	GrpcServer *server.Server
	statusOnce = &sync.Once{}
	Mregister  = &serverplugin.NacosRegisterPlugin{}
)

//初始化grpc service服务
func StartGrpcService(serviceName, rpcAddr, port string, nacosAddrs []string) {
	statusOnce.Do(func() {
		GrpcServer = server.NewServer()
		addNacosRegistryPlugin(nacosAddrs, port)
		if err := Mregister.Register(serviceName, &LarkImpl{}, ""); err != nil {
			panic(err)
		}
		GrpcServer.RegisterName(serviceName, &LarkImpl{}, "")
		if err := GrpcServer.Serve("tcp", rpcAddr); err != nil {
			panic(err)
		}
	})
}

func addNacosRegistryPlugin(nacosAddrs []string, port string) {
	serverConfigs := utils.GetRegisterServerConf(nacosAddrs)
	Mregister = &serverplugin.NacosRegisterPlugin{
		ServiceAddress: "tcp@" + utils.ExternalIP() + ":" + port,
		ServerConfig:   serverConfigs,
	}
	err := Mregister.Start()
	if err != nil {
		panic(err)
	}
	GrpcServer.Plugins.Add(Mregister)
}
