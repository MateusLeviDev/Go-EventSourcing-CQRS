package interceptors

import (
	"context"
	"time"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)

	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
}

// InterceptorManager struct
type interceptorManager struct {
	log logger.Logger
}

// NewInterceptorManager InterceptorManager constructor
func NewInterceptorManager(logger logger.Logger) *interceptorManager {
	return &interceptorManager{log: logger}
}

// Logger Interceptor
func (im *interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.log.GrpcMiddlewareAccessLogger(info.FullMethod, time.Since(start), md, err)
	return reply, err

}

// ClientRequestLoggerInterceptor gRPC client interceptor
func (im *interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		md, _ := metadata.FromIncomingContext(ctx)
		im.log.GrpcClientInterceptorLogger(method, req, reply, time.Since(start), md, err)
		return err
	}

}
