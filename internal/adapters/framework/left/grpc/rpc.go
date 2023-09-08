package rpc

import (
	"context"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ad *Adapter) GetAddition(ctx context.Context, req *pb.Params) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	ans, err := ad.api.GetAddition(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: ans}, nil
}

func (ad *Adapter) GetSubtraction(ctx context.Context, req *pb.Params) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	ans, err := ad.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: ans}, nil
}

func (ad *Adapter) GetMultiplication(ctx context.Context, req *pb.Params) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	ans, err := ad.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: ans}, nil
}

func (ad *Adapter) GetDivision(ctx context.Context, req *pb.Params) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	ans, err := ad.api.GetDivision(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: ans}, nil
}
