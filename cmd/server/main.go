package main

import (
	"ent-study/adapter/in"
	pb "ent-study/proto/v1/member"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("filaed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterMemberServiceServer(server, in.NewMemberGrpcAdapter())
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
