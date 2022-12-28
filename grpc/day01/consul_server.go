package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	//"github.com/hashicorp/consul"
	"github.com/hashicorp/consul/api"
	"grpc/day01/pb"
)

type Children struct {
}

// 绑定类方法实现接口
func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello" + p.Name
	return p, nil
}
func main() {

	/*grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	listener, err := net.Listen("tcp", "127.0.0.1:8800")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer listener.Close()

	// 启动服务
	grpcServer.Serve(listener)*/

	//初始化配置
	consulConfig := api.DefaultConfig()

	//创建consul对象
	consulClient, error := api.NewClient(consulConfig)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	// 告诉consul即将注册的服务信息配置
	registerService := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{"testhello"},
		Name:    "HelloService",
		Port:    8800,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8800",
			Timeout:  "5s",
			Interval: "5s",
		},
	}

	consulClient.Agent().ServiceRegister(&registerService)

	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	listener, err := net.Listen("tcp", "127.0.0.1:8800")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer listener.Close()

	fmt.Println("启动成功")
	// 启动服务
	grpcServer.Serve(listener)

}
