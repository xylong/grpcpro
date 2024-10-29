package services

import (
	"context"
	services "grpcpro/pbfiles"
)

type ProdService struct {
	services.UnimplementedProdServiceServer
}

func (s *ProdService) GetProdStock(ctx context.Context, in *services.ProdRequest) (*services.ProdResponse, error) {
	return &services.ProdResponse{ProdStock: 25}, nil
}

//func (s *ProdService) mustEmbedUnimplementedProdServiceServer() {
//
//}
