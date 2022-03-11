package todolist

import (
	"context"
	"fmt"
	"log"
	"main/todolist/infra/mongo"
	pb "main/todolist/protobuf"
	todolist "main/todolist/protobuf"
	"main/todolist/srv"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

// Create - 新增事項
func (*Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	name := req.GetName()
	memberId := req.GetMemberId()

	res := &pb.CreateResponse{}

	if err := srv.Create(ctx, name, memberId); err != nil {
		return res, err
	}

	return res, nil
}

// Delete - 刪除事項
func (*Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	id := req.GetID()
	memberId := req.GetMemberId()

	res := &pb.DeleteResponse{}

	if err := srv.Delete(ctx, id, memberId); err != nil {
		return res, err
	}

	return res, nil
}

// Get - 取得事項
func (*Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	memberId := req.GetMemberId()

	res := &pb.GetResponse{}
	results, err := srv.Get(ctx, memberId)
	if err != nil {
		return res, err
	}

	return &pb.GetResponse{
		TodoList: todolist.ConverTodoListToPb(results),
	}, nil
}

// Update - 完成事項
func (*Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	id := req.GetID()
	memberId := req.GetMemberId()

	res := &pb.UpdateResponse{}
	err := srv.Update(ctx, id, memberId)
	if err != nil {
		return res, err
	}

	return res, nil
}

// init - 初始化
func init() {
	mongo.Init()
}

// TodoServer - 待辦事項服務
func TodoServer() {
	fmt.Println("starting gRPC todo server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
