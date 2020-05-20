package main

import (
	"crypto/tls"
	"crypto/x509"
	"go-grpc-server/service"
	services "go-grpc-server/service/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//创建一个RPC服务器
	//rpcServer := grpc.NewServer()
	rpcServer := grpc.NewServer(grpc.Creds(GetServerCreds()))

	//在RPC服务器中注册服务
	services.RegisterHelloServiceServer(rpcServer,new(service.HelloService))
	services.RegisterProductServiceServer(rpcServer,new(service.ProductService))
	//定义一个TCP监听器
	listener, _ := net.Listen("tcp", ":8081")

	//RPC服务器在指定TCP监听器监听客户连接
	rpcServer.Serve(listener)
}

func GetServerCreds() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},   // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}
