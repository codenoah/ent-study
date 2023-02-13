package main

import (
	"context"
	pb "ent-study/proto/v1/member"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.JoinMember(ctx, &pb.JoinMember_Request{AccountId: "test", Password: "조아"})
	if err != nil {
		log.Fatalf("could not join member: %v", err)
	}
	log.Println("join Member response: ", res)
}
