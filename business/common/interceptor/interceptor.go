package interceptor

import (
	"context"
	"github.com/dubbogo/gost/log/logger"
	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info("RPC---该请求为", info.FullMethod)
	resp, err := handler(ctx, req)
	logger.Info("RPC---该请求返回", resp, err)
	return resp, err
}
