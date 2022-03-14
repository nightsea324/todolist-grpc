package member

import (
	"context"
	"fmt"
	"log"
	"main/member/infra/mongo"
	"main/member/infra/redis"
	pb "main/member/protobuf"
	"main/member/srv"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

// Register - 註冊會員
func (*Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	fmt.Println("Register...")
	name := req.GetName()
	password := req.GetPassword()

	res := &pb.RegisterResponse{}

	if err := srv.Register(ctx, name, password); err != nil {
		return res, err
	}

	return res, nil
}

// SignIn - 登入會員
func (*Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {

	fmt.Println("SignIn...")
	name := req.GetName()
	password := req.GetPassword()

	res := &pb.SignInResponse{}
	token, tokenLife, err := srv.SignIn(ctx, name, password)
	if err != nil {
		return res, err
	}

	return &pb.SignInResponse{
		Token:     token,
		TokenLife: tokenLife}, nil
}

// init - 初始化
func init() {
	mongo.Init()
	redis.Init()
}

// MemberServer - 會員服務
func MemberServer() {
	fmt.Println("starting gRPC member server...")

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
