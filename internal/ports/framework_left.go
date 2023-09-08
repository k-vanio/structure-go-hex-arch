package ports

import (
	"context"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc/pb"
)

type GRPC interface {
	Run()
	GetAddition(context.Context, *pb.Params) (*pb.Answer, error)
	GetSubtraction(context.Context, *pb.Params) (*pb.Answer, error)
	GetMultiplication(context.Context, *pb.Params) (*pb.Answer, error)
	GetDivision(context.Context, *pb.Params) (*pb.Answer, error)
}
