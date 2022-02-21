// Package middleware contains grpc validation and authorization interceptors
package middleware

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/golang-jwt/jwt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func authFunc(ctx context.Context) (*config.TokenClaims, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no auth meta-data found in request")
	}
	mapValue := meta["authorization"][0]
	var token string
	_, err := fmt.Sscanf(mapValue, "Bearer %s", &token)
	if err != nil {
		return nil, fmt.Errorf("scan error in authFunc in interceptor %v", err)
	}
	claims, err := validateToken(token)
	if err != nil {
		return nil, fmt.Errorf("futh func in inerceptor %v", err)
	}
	return claims, nil
}

func validateToken(token string) (*config.TokenClaims, error) {
	tokenType, err := jwt.ParseWithClaims(token, &config.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.AccessKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token error %v", err)
	}
	if claims, ok := tokenType.Claims.(*config.TokenClaims); ok && !tokenType.Valid {
		return nil, fmt.Errorf("validate token func %v", err)
	} else {
		return claims, nil
	}
}
func unaryServerAuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// skip authorization in log out and refresh
	if info.FullMethod != "/proto.|||Service/LogOut" && info.FullMethod != "/proto.|||Service/Refresh" {
		if _, err := authFunc(ctx); err != nil {
			return nil, err
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}

func streamServerAuthorizationInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if _, err := authFunc(ss.Context()); err != nil {
		return fmt.Errorf("auth func error in stream server interceptor %v", err)
	}
	wrapped := grpc_middleware.WrapServerStream(ss)
	return handler(srv, wrapped)
}
