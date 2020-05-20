package service

import (
	"context"
	"fmt"
	proxy "go-grpc-server/service/api"
)

type ProductService struct {
}

func (this *ProductService) GetProductById(ctx context.Context, request *proxy.ProductRequest) (*proxy.ProductResponse, error) {
	fmt.Println("客户请求: GetProductById 方法")
	return &proxy.ProductResponse{Name: "Name" + request.Id, Code: "Code" + request.Id}, nil
}

func (this *ProductService) GetProductList(ctx context.Context, request *proxy.QuerySize) (*proxy.ProductList, error) {
	productList := make([]*proxy.ProductResponse,int(request.Size))
	for i := 0; i < int(request.Size); i++ {
		productList[i]=&proxy.ProductResponse{Name: fmt.Sprintf("Name%d",i),Code: fmt.Sprintf("Code%d",i)}
		//append(productList, &ProductResponse{Name: "Name"+string(i),Code: "Code"+string(i)})
	}
	fmt.Println("客户请求: GetProductList 方法")
	return &proxy.ProductList{ProductList: productList},nil
}
