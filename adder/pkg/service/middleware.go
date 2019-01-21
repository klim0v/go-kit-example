package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AdderService) AdderService

type loggingMiddleware struct {
	logger log.Logger
	next   AdderService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AdderService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AdderService) AdderService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Sum(ctx context.Context, a int, b int) (rs int, err error) {
	defer func() {
		l.logger.Log("method", "Sum", "a", a, "b", b, "rs", rs, "err", err)
	}()
	return l.next.Sum(ctx, a, b)
}
