package logging

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

//MiddlewareEndpoint logs the endpoint calling
func MiddlewareEndpoint(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			begin := time.Now()
			logger.Log("message", "calling endpoint")
			defer logger.Log(
				"message", "end of the endpoint calling",
				"took", time.Since(begin),
			)
			return next(ctx, request)
		}
	}
}
