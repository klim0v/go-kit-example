package service

import "context"

// AdderService describes the service.
type AdderService interface {
	// Add your methods here
	Sum(ctx context.Context, a, b int) (rs int, err error)
}

type basicAdderService struct{}

func (ba *basicAdderService) Sum(ctx context.Context, a int, b int) (rs int, err error) {
	// TODO implement the business logic of Sum
	return rs, err
}

// NewBasicAdderService returns a naive, stateless implementation of AdderService.
func NewBasicAdderService() AdderService {
	return &basicAdderService{}
}

// New returns a AdderService with all of the expected middleware wired in.
func New(middleware []Middleware) AdderService {
	var svc AdderService = NewBasicAdderService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
