package main

import (
	grpc2 "github.com/HekapOo-hub/Task1/internal/grpc"
	"github.com/HekapOo-hub/Task1/internal/middleware"
	pb "github.com/HekapOo-hub/Task1/internal/proto/humanpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

func main() {

	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("error %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
		grpc.StreamInterceptor(middleware.StreamServerInterceptor()),
	)
	handler, cancel, err := grpc2.GetInstance()
	if err != nil {
		log.Warnf("%v", err)
		return
	}
	defer cancel()
	pb.RegisterHumanServiceServer(s, handler)
	go func() {
		time.Sleep(time.Second * 15)
		os.Exit(123)
	}()
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
