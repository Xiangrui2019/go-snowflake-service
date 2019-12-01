package services

import (
	"context"
	"go-snowflake-service/generator_pool"
	"go-snowflake-service/pb"
)

type IdGeneratorService struct {
}

func (service *IdGeneratorService) GenerateId(ctx context.Context, in *pb.GenerateIdRequest) (*pb.GenerateIdReply, error) {
	flake := generator_pool.GetSnowFlake(in.NodeId)

	id, err := flake.GenerateId()

	return &pb.GenerateIdReply{Id: int64(id)}, err
}
