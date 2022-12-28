package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"grpc/day01/pb"
	"strconv"
)

// 客户端调用grpc服务
func main() {

	// 1.初始化consul配置
	consulConfig := api.DefaultConfig()

	//2.创建consul对象
	consulClient, error := api.NewClient(consulConfig)

	// 3. 从consul上获取健康的服务

	ServiceEntry, _, error := consulClient.Health().Service("HelloService", "testhello", true, nil)

	if error != nil {

		fmt.Println(error.Error())
		fmt.Println("获取服务失败")
		return

	}

	address := ServiceEntry[0].Service.Address + ":" + strconv.Itoa(ServiceEntry[0].Service.Port)
	//grpcconn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	grpcconn, _ := grpc.Dial(address, grpc.WithInsecure())

	//初始化grpc客户端

	grpcClient := pb.NewHelloClient(grpcconn)

	var person pb.Person
	person.Name = "judy"
	person.Age = 10
	//调用远程函数
	p, error := grpcClient.SayHello(context.TODO(), &person)

	fmt.Println(p, error)
}
