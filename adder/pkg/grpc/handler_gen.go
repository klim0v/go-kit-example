// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/klim0v/go-kit-example/adder/pkg/endpoint"
	pb "github.com/klim0v/go-kit-example/adder/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sum grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.AdderServer {
	return &grpcServer{sum: makeSumHandler(endpoints, options["Sum"])}
}
