package grpc

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/akabos/ctxlog"
)

func UnaryWithLogger(l zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(ctxlog.WithLogger(ctx, l), req)
	}
}

func StreamWithLogger(l zerolog.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, &serverStream{stream, ctxlog.WithLogger(stream.Context(), l)})
	}
}

func UnaryWithLoggerFrom(ctx context.Context) grpc.UnaryServerInterceptor {
	return UnaryWithLogger(ctxlog.Logger(ctx))
}

func StreamWithLoggerFrom(ctx context.Context) grpc.StreamServerInterceptor {
	return StreamWithLogger(ctxlog.Logger(ctx))
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *serverStream) Context() context.Context {
	return s.ctx
}
