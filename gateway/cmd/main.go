package main

import (
	"fmt"
	"gateway/common/discovery"
	"gateway/common/global"
	_ "gateway/common/init"
	"gateway/middleware/wrapper"
	service "gateway/proto"
	"gateway/routers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go startListen()
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	config := global.Config
	// etcd 注册
	etcdRegister := discovery.NewResolver([]string{config.Etcd.Address}, logrus.New())
	resolver.Register(etcdRegister)
	//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// 服务
	userServiceName := config.Domain.User
	taskServiceName := config.Domain.Task

	// GRPC connetcion
	userConnect, err := GRPCConnect(userServiceName, etcdRegister)
	if err != nil {
		return
	}
	userClient := service.NewUserServiceClient(userConnect)
	taskConnect, err := GRPCConnect(taskServiceName, etcdRegister)
	if err != nil {
		return
	}
	taskClient := service.NewTaskServiceClient(taskConnect)

	// 熔断
	wrapper.NewServiceWrapper(userServiceName)
	wrapper.NewServiceWrapper(taskServiceName)

	routers.HttpServerRun([]interface{}{userClient, taskClient})
}

func GRPCConnect(serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)

	return grpc.Dial(addr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	}...)
}
