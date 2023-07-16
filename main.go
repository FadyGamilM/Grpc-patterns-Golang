package main

import (
	"fmt"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello")
	logger := hclog.Default()

	gs := grpc.NewServer()

}