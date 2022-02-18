package middleware

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/validation"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func unaryServerValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	validator, err := validation.NewValidator()
	if err != nil {
		return nil, fmt.Errorf("creater validator error %v", err)
	}
	err = validator.Validate(req)
	if err != nil {
		return nil, fmt.Errorf("unary server validation interceptor error %v", err)
	}

	h, err := handler(ctx, req)

	return h, err
}

func streamServerValidationInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	validator, err := validation.NewValidator()
	if err != nil {
		return fmt.Errorf("creater validator error %v", err)
	}
	var req interface{}
	err = ss.RecvMsg(req)
	err = validator.Validate(req)
	if err != nil {
		return fmt.Errorf("stream server validation interceptor error %v", err)
	}
	wrapped := grpc_middleware.WrapServerStream(ss)
	return handler(srv, wrapped)
}

func WithUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(unaryServerValidationInterceptor, unaryServerAuthorizationInterceptor)
}

func WithStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(streamServerValidationInterceptor, streamServerAuthorizationInterceptor)
}
