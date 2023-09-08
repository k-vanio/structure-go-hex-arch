package rpc

import (
	"log"
	"net"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc/pb"
	"github.com/k-vanio/structure-go-hex-arch/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.Api
}

func NewAdapter(api ports.Api) *Adapter {
	return &Adapter{api: api}
}

func (ad *Adapter) Run() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}

	srv := ad
	server := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(server, srv)

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
