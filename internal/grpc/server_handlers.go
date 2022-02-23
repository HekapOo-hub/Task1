package grpc

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/middleware"
	"github.com/HekapOo-hub/Task1/internal/model"
	pb "github.com/HekapOo-hub/Task1/internal/proto/humanpb"
	"github.com/HekapOo-hub/Task1/internal/service"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

const (
	admin = "admin"
)

type HumanServer struct {
	pb.UnimplementedHumanServiceServer
	humanService *service.HumanService
	userService  *service.UserService
	authService  *service.AuthService
	fileService  *service.FileService
}

func (s *HumanServer) CreateHuman(ctx context.Context, in *pb.Human) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("crate human, server error %v", err)
		return nil, fmt.Errorf("create human, server error %w", err)
	}
	role := claims.Role
	if role != admin {
		log.Warn("access denied")
		return nil, fmt.Errorf("access denied")
	}

	err = s.humanService.Create(ctx, model.Human{Name: in.Name, Male: in.Male, Age: int(in.Age)})
	if err != nil {
		log.Warnf("error: %v", err)
		return nil, fmt.Errorf("create human, server error %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) UpdateHuman(ctx context.Context, in *pb.UpdateHumanRequest) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("update human, server error %v", err)
		return nil, fmt.Errorf("update human, server error %w", err)
	}
	if claims.Role != admin {
		log.Warn("access denied")
		return nil, fmt.Errorf("access denied")
	}

	err = s.humanService.Update(ctx, in.OldName, model.Human{Name: in.Human.Name,
		Male: in.Human.Male, Age: int(in.Human.Age)})
	if err != nil {
		log.Warnf("update human, server error %v", err)
		return nil, fmt.Errorf("update human, server error %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) GetHuman(ctx context.Context, in *pb.Name) (*pb.Human, error) {
	human, err := s.humanService.Get(ctx, in.Value)
	if err != nil {
		log.Warnf("error: %v", err)
		return nil, fmt.Errorf("get human, server error %w", err)
	}
	return &pb.Human{ID: human.ID, Name: human.Name, Age: int32(human.Age), Male: human.Male}, nil
}

func (s *HumanServer) DeleteHuman(ctx context.Context, in *pb.Name) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("delete human, server error %v", err)
		return nil, fmt.Errorf("delete human, server error %w", err)
	}
	if claims.Role != admin {
		log.Warn("access denied")
		return nil, fmt.Errorf("access denied")
	}
	err = s.humanService.Delete(ctx, in.Value)
	if err != nil {
		log.Warnf("delete human, server error: %v", err)
		return nil, fmt.Errorf("delete human, server error %v", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) Authenticate(ctx context.Context, in *pb.SignInRequest) (*pb.Tokens, error) {
	user, err := s.userService.Get(ctx, in.Login)
	if err != nil {
		log.Warnf("grpc authenticate: %v", err)
		return nil, fmt.Errorf("grpc authenticate %w", err)
	}
	accessToken, refreshToken, err := s.authService.Authenticate(ctx, user, in.Password)
	if err != nil {
		log.Warnf("error with token in authentication: %v", err)
		return nil, fmt.Errorf("grpc authenticate %w", err)
	}
	return &pb.Tokens{Access: accessToken, Refresh: refreshToken}, nil
}

func (s *HumanServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc create user: %v", err)
		return nil, fmt.Errorf("grpc create user %w", err)
	}
	if claims.Role != admin {
		return nil, fmt.Errorf("access denied in create user")
	}
	err = s.userService.Create(ctx, in.Login, in.Password)
	if err != nil {
		log.Warnf("grpc creating user: %v", err)
		return nil, fmt.Errorf("grpc create user %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) GetUser(ctx context.Context, in *pb.Login) (*pb.User, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc get user: %v", err)
		return nil, fmt.Errorf("grpc get user %w", err)
	}
	if in.Value != claims.Login && claims.Role != admin {
		return nil, fmt.Errorf("access denied in grpc get user")
	}
	user, err := s.userService.Get(ctx, in.Value)
	if err != nil {
		log.Warnf("grpc get user %v", err)
		return nil, fmt.Errorf("grpc get user %w", err)
	}
	return &pb.User{Login: user.Login, ID: user.ID, Role: user.Role}, nil
}

func (s *HumanServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc update user: %v", err)
		return nil, fmt.Errorf("grpc update user %w", err)
	}
	if claims.Login != in.OldLogin && claims.Role != admin {
		log.Warnf("access denied in grpc update user")
		return nil, fmt.Errorf("access denied in grpc update user")
	}
	err = s.userService.Update(ctx, in.OldLogin,
		model.User{Login: in.NewLogin, Password: in.NewPassword})
	if err != nil {
		log.Warnf("grpc update user %v", err)
		return nil, fmt.Errorf("grpc update user %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) DeleteUser(ctx context.Context, in *pb.Login) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc delete user: %v", err)
		return nil, fmt.Errorf("grpc delete user %w", err)
	}
	if claims.Login != in.Value && claims.Role != admin {
		log.Warnf("access denied in grpc delete user")
		return nil, fmt.Errorf("access denied in grpc delete user")
	}
	err = s.userService.Delete(ctx, in.Value)
	if err != nil {
		log.Warnf("access denied in grpc delete user")
		return nil, fmt.Errorf("access denied in grpc delete user")
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) Refresh(ctx context.Context, in *pb.Tokens) (*pb.Tokens, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc delete user: %v", err)
		return nil, fmt.Errorf("grpc delete user %w", err)
	}
	accessToken, refreshToken, err := s.authService.Refresh(ctx, claims, in.Refresh)
	if err != nil {
		log.Warnf("grpc refresh %v", err)
		return nil, fmt.Errorf("grpc refresh %w", err)
	}
	return &pb.Tokens{Access: accessToken, Refresh: refreshToken}, nil
}

func (s *HumanServer) LogOut(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	claims, err := middleware.AuthFunc(ctx)
	if err != nil {
		log.Warnf("grpc delete user: %v", err)
		return nil, fmt.Errorf("grpc delete user %w", err)
	}
	err = s.authService.Delete(ctx, claims.Login)
	if err != nil {
		log.Warnf("grpc log out %v", err)
		return nil, fmt.Errorf("grpc log out %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *HumanServer) DownloadFile(stream pb.HumanService_DownloadFileServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Warnf("grpc stream download file %v", err)
				return fmt.Errorf("grpc stream download file %w", err)
			}
			return nil
		default:
		}
		start := true
		offset := 0
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("receive error in download %v", err)
			continue
		}
		file, err := s.fileService.Download(ctx, req.Value)

		if err != nil {
			log.Warnf("grpc download file %v", err)
			return fmt.Errorf("grpc download file %v", err)
		}

		for {
			buff := make([]byte, 400)

			if _, err = file.ReadAt(buff, int64(offset)); err == nil || err == io.EOF {
				if streamErr := stream.Send(&pb.FilePortion{Value: buff, Start: start}); streamErr != nil {
					log.Warnf("sending file portion to stream error in download %v", streamErr)
					return fmt.Errorf("sending file portion to stream error in download %w", streamErr)
				}
			} else {
				log.Warnf("reading file in download grpc %v", err)
				return fmt.Errorf("reading file in download grpc %v", err)
			}
			if err == io.EOF {
				break
			}
			offset += 400
			start = false
		}
	}
}

func (s *HumanServer) UploadFile(stream pb.HumanService_UploadFileServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Warnf("grpc stream upload file %v", err)
				return fmt.Errorf("grpc stream upload file %w", err)
			}
		default:
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("exit from upload stream")
			return nil
		}
		if err != nil {
			log.Printf("receive error in upload %v", err)
			continue
		}
		file, err := os.Open(filepath.Clean(req.Value))
		if err != nil {
			return fmt.Errorf("open file error in upload %w", err)
		}

		err = s.fileService.Upload(ctx, file)
		if err != nil {
			log.Warnf("grpc stream upload %v", err)
			return fmt.Errorf("grpc stream upload %w", err)
		}
		if err := stream.Send(&pb.Empty{}); err != nil {
			log.Warnf("send error in upload %v", err)
			return fmt.Errorf("send error in upload %v", err)
		}
	}
}
