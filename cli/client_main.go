package main

import (
	"context"
	pb "github.com/HekapOo-hub/Task1/internal/proto/humanpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := pb.NewHumanServiceClient(conn)
	signInRequest := pb.SignInRequest{Login: "admin", Password: "1234"}
	tokens, err := client.Authenticate(context.Background(), &signInRequest)
	if err != nil {
		log.Warnf("%v", err)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": "Bearer " + tokens.Access}))

	_, err = client.GetHuman(ctx, &pb.Name{Value: "me"})
	if err != nil {
		log.Warnf("%v", err)
	}

	uploadStream, err := client.UploadFile(ctx)
	if err != nil {
		log.Warnf("stream open error %v", err)
	}
	go func() {
		for i := 1; i < 5; i++ {
			if err := uploadStream.Send(&pb.Name{Value: "img.png"}); err != nil {
				log.Warnf("send error %v", err)
			}
			time.Sleep(time.Millisecond * 200)
		}
		time.Sleep(time.Second * 2)
		_ = uploadStream.CloseSend()
	}()
	done := make(chan bool)
	go func() {
		for {
			resp, err := uploadStream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			if resp == nil {
				log.Warnf("error with stream upload func in recieving!!!")
			}
		}
	}()
	time.Sleep(time.Second)
	downloadStream, err := client.DownloadFile(ctx)
	if err != nil {
		log.Warnf("%v", err)
	}
	go func() {
		for i := 1; i < 5; i++ {
			if err := downloadStream.Send(&pb.Name{Value: "img.png"}); err != nil {
				log.Warnf("send error %v", err)
			}
			time.Sleep(time.Millisecond * 500)
		}
		_ = downloadStream.CloseSend()
	}()
	counter := 0
	var file *os.File
	go func() {
		for {
			resp, err := downloadStream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			if resp.Start {
				counter++
				log.Info("create new file")
				file, err = os.Create("cpy" + strconv.Itoa(counter) + ".png")
				if err != nil {
					log.Warnf("%v", err)
					return
				}
			}
			_, err = file.Write(resp.Value)
			if err != nil {
				log.Warnf("%v", err)
			}
		}
	}()
	<-done
}
