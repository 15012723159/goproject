package main

import (
	"context"
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
	api.DefaultConfig
}
