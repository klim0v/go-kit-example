package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/klim0v/go-kit-example/adder/pkg/endpoint"
	"github.com/klim0v/go-kit-example/adder/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeSumHandler creates the handler logic
func makeSumHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SumEndpoint, decodeSumRequest, encodeSumResponse, options...)
}

// decodeSumResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeSumRequest(_ context.Context, r interface{}) (interface{}, error) {
	request := r.(*pb.SumRequest)
	return endpoint.SumRequest{A: request.A, B: request.B}, nil
}

// encodeSumResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSumResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.SumResponse)
	return &pb.SumReply{Rs: response.Rs}, nil
}
func (g *grpcServer) Sum(ctx context1.Context, req *pb.SumRequest) (*pb.SumReply, error) {
	_, rep, err := g.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SumReply), nil
}
