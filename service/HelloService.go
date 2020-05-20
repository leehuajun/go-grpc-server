package service

import (
	"context"
	"fmt"
	. "go-grpc-server/service/api"
)

type HelloService struct {

}

func (this *HelloService) SayHello(ctx context.Context, request *HelloRequest)(*HelloResponse,error){
	fmt.Println("客户请求: SayHello 方法")
	response:= HelloResponse{Message: "Hello,"+request.Name}
	return &response,nil

}