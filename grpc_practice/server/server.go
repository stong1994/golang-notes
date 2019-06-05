package main

import (
	"context"
	"github.com/pkg/errors"
	pb "golang-learning/grpc_practice/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServer struct {
}

func (us *UserServer) Login(ctx context.Context, u *pb.LoginRequest) (*pb.LoginResponse, error) {
	if u.GetAge() == 18 && u.GetName() == "andy" {
		return &pb.LoginResponse{Success: true}, nil
	}
	return nil, errors.New("can not access auth")
}

func main() {
	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	opt := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opt...)

	pb.RegisterUserServerServer(grpcServer, &UserServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
