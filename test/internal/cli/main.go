package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcpro/pbfiles"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := pbfiles.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &pbfiles.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====")
	fmt.Println(prodRes.ProdStock)
}
