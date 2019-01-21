package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/klim0v/go-kit-example/adder/pkg/service"
)

// SumRequest collects the request parameters for the Sum method.
type SumRequest struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

// SumResponse collects the response parameters for the Sum method.
type SumResponse struct {
	Rs  int32 `json:"rs"`
	Err error `json:"err"`
}

// MakeSumEndpoint returns an endpoint that invokes Sum on the service.
func MakeSumEndpoint(s service.AdderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		rs, err := s.Sum(ctx, req.A, req.B)
		return SumResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SumResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Sum implements Service. Primarily useful in a client.
func (e Endpoints) Sum(ctx context.Context, a int32, b int32) (rs int32, err error) {
	request := SumRequest{
		A: a,
		B: b,
	}
	response, err := e.SumEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SumResponse).Rs, response.(SumResponse).Err
}
