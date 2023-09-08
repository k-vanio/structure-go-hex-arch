package rpc_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/app/api"
	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/core/arithmetic"
	rpc "github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc"
	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc/pb"
	"github.com/k-vanio/structure-go-hex-arch/internal/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context, dbAdapter ports.Db) (pb.ArithmeticServiceClient, func()) {
	// ports
	var coreAdapter ports.Arithmetic
	var appAdapter ports.Api
	var gRPCAdapter ports.GRPC

	coreAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(coreAdapter, dbAdapter)

	gRPCAdapter = rpc.NewAdapter(appAdapter)

	buffer := 1024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(baseServer, gRPCAdapter)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewArithmeticServiceClient(conn)

	return client, closer
}

type DBStub struct {
	mock.Mock
}

func (db *DBStub) Close() {}
func (db *DBStub) AddHistory(answer int32, operation string) error {
	args := db.Called(answer, operation)
	return args.Error(0)
}

func TestGetAdditionSuccessfully(t *testing.T) {
	ctx := context.Background()
	dbStub := new(DBStub)

	client, closer := server(ctx, dbStub)
	defer closer()

	dbStub.On("AddHistory", int32(7), "Addition").Return(nil)

	an, err := client.GetAddition(ctx, &pb.Params{A: 3, B: 4})
	expected := &pb.Answer{Value: 7}

	assert.Nil(t, err)
	assert.Equal(t, expected.GetValue(), an.GetValue())
	dbStub.AssertNumberOfCalls(t, "AddHistory", 1)
	dbStub.MethodCalled("AddHistory", int32(7), "Addition")
	dbStub.AssertExpectations(t)
}
