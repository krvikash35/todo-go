package server

import (
	"net"

	"google.golang.org/grpc"
)

func StartGRPC() {

	l, _ := net.Listen("tcp", "8082")
	s := grpc.NewServer()
	s.Serve(l)
}
