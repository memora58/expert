package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"user/common/discovery"
	"user/common/global"
	_ "user/common/init"
	"user/controller"
	service "user/proto"
)

func main() {
	config := global.Config
	// 服务注册
	etcdRegister := discovery.NewRegister([]string{config.Etcd.Address}, logrus.New())
	defer etcdRegister.Stop()
	taskNode := discovery.Server{
		Name: config.Server.Domain,
		Addr: config.Server.GrpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定 task service
	service.RegisterUserServiceServer(server, &controller.UserController{})
	lis, err := net.Listen("tcp", config.Server.GrpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(taskNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", config.Server.GrpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
