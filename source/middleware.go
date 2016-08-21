package source

import (
	"golang.org/x/net/context"
	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next Service
	logger log.Logger
}

func (mw *loggingMiddleware) PostSource(ctx context.Context, s *Source) error {
	mw.logger.Log("method", "PostSource")
	return mw.next.PostSource(ctx, s)
}

func (mw *loggingMiddleware) GetSource(ctx context.Context, uuidStr string) (Source, error) {
	mw.logger.Log("method", "GetSource")
	return mw.next.GetSource(ctx, uuidStr)
}

func (mw *loggingMiddleware) PutSource(ctx context.Context, uuidStr string, s *Source) error {
	mw.logger.Log("method", "PutSource")
	return mw.next.PutSource(ctx, uuidStr, s)
}

func (mw *loggingMiddleware) PatchSource(ctx context.Context, uuidStr string, s *Source) error {
	mw.logger.Log("method", "PatchSource")
	return mw.next.PatchSource(ctx, uuidStr, s)
}

func (mw *loggingMiddleware) DeleteSource(ctx context.Context, uuidStr string) error {
	mw.logger.Log("method", "DeleteSource")
	return mw.next.DeleteSource(ctx, uuidStr)
}

func (mw *loggingMiddleware) GetSources(ctx context.Context) ([]Source, error) {
	mw.logger.Log("method", "GetSources")
	return mw.next.GetSources(ctx)
}