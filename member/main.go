package main

import (
	"context"
	"fmt"
	"log"
	"main/mongo"
	pb "main/protobuf/member"
	"main/srv"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	name := req.GetName()
	password := req.GetPassword()
	var status string
	var msg string

	msg, err := srv.Register(name, password)
	if err != nil {
		status = "failed"
	} else {
		status = "ok"
		msg = "註冊成功"
	}

	res := &pb.RegisterResponse{
		Status: status,
		Msg:    msg,
	}

	return res, nil
}

func (*Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {

	res := &pb.SignInResponse{
		Status: "ok",
		Msg:    "註冊成功",
	}

	return res, nil
}

func main() {
	fmt.Println("starting gRPC server...")
	mongo.Init()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMemberServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
