package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/FadyGamilM/gogrpc/protos/currency"
	"github.com/FadyGamilM/gogrpc/server"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("hello")
	logger := hclog.Default()

	gs := grpc.NewServer()

	reflection.Register(gs)

	cs:= server.NewCurrencyServer(logger)

	protos.RegisterCurrencyServer(gs, cs)
	// start the grpc server 

	listener, err := net.Listen("tcp", ":5050")
	if err!= nil {
		logger.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)

}

// to use the grpcurl to list all the services in our system use this command
// > grpcurl --plaintext localhost:5050 list

// to use the grpcurl to list all the methods inside the service in our system use this command
// > grpcurl --plaintext localhostom :5050 list Currency

/*
➜ Grpc-patterns-Golang git:(main) grpcurl --plaintext localhost:5050 describe Currency
	Currency is a service:
	service Currency {
	rpc GetRate ( .RateReq ) returns ( .RateRes );
	}
*/