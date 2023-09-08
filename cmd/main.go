package main

import (
	"log"
	"os"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/app/api"
	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/core/arithmetic"
	rpc "github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/left/grpc"
	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/framework/right/db"
	"github.com/k-vanio/structure-go-hex-arch/internal/ports"
)

func main() {
	// ports
	var dbAdapter ports.Db
	var coreAdapter ports.Arithmetic
	var appAdapter ports.Api
	var gRPCAdapter ports.GRPC

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	dbAdapter, err := db.NewAdapter(dbDriver, dbSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbAdapter.Close()

	coreAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(coreAdapter, dbAdapter)

	gRPCAdapter = rpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
