package main

import (
	"google.golang.org/grpc"
	"grpcpro/pbfiles"
	"grpcpro/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	lis, _ := net.Listen("tcp", ":8081")

	rpcServer.Serve(lis)

}
