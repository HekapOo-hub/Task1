package middleware

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/validation"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
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

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	validator, err := validation.NewValidator()
	if err != nil {
		return fmt.Errorf("creater validator error %v", err)
	}

	if err := validator.Validate(m); err != nil {
		log.Warnf("stream server validation interceptor error %v", err)
		return fmt.Errorf("stream server validation interceptor error %v", err)
	}

	return nil
}

func streamServerValidationInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	wrapper := &recvWrapper{ss}
	return handler(srv, wrapper)
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(unaryServerValidationInterceptor, unaryServerAuthorizationInterceptor)
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(
		streamServerValidationInterceptor,
		streamServerAuthorizationInterceptor,
	)
}
