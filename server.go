package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcpro/pbfiles"
	"grpcpro/services"
	"net/http"
)

func main() {

	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//rpcServer:=grpc.NewServer(grpc.Creds(creds))
	rpcServer := grpc.NewServer()
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	//lis, _ := net.Listen("tcp", ":8081")
	//rpcServer.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	httpServer.ListenAndServe()
	//httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_passwd.key")
}
