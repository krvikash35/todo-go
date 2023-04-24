package grpc

import (
	context "context"
	"log"
	"net"

	grpcpkg "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	UnimplementedToDoServiceServer
}

func (s *server) Ping(ctx context.Context, in *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: "pong"}, nil
}

func StartServer() {

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("StartGRPC: failed to listen: %v", err)
	}

	s := grpcpkg.NewServer()
	reflection.Register(s)
	RegisterToDoServiceServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatalf("StartGRPC: failed to serve: %v", err)
	}
}
