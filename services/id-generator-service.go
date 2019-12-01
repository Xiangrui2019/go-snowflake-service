package services

import (
	"context"
	"go-snowflake-service/pb"
)

type IdGeneratorService struct {
}

func (service *IdGeneratorService) GenerateId(ctx context.Context, in *pb.GenerateIdRequest) (*pb.GenerateIdReply, error) {
	return &pb.GenerateIdReply{
		Id:    222,
		Error: "",
	}, nil
}
