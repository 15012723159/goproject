package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/day01/pb"
)

// 客户端调用grpc服务
func main() {
	//
	grpcconn, _ := grpc.Dial(":8800", grpc.WithInsecure())

	//初始化grpc客户端

	grpcClient := pb.NewHelloClient(grpcconn)

	var person pb.Person
	person.Name = "judy"
	person.Age = 10
	//调用远程函数
	p, error := grpcClient.SayHello(context.TODO(), &person)

	fmt.Println(p, error)
}
