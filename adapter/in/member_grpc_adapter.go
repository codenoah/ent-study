package in

import (
	"context"
	"ent-study/application/port/input"
	"ent-study/application/service"
	pb "ent-study/proto/v1/member"
	"log"
)

type memberGrpcAdapter struct {
	pb.UnimplementedMemberServiceServer
	memberService input.MemberPort
}

func NewMemberGrpcAdapter() *memberGrpcAdapter {
	return &memberGrpcAdapter{memberService: service.NewMemberService()}
}

func (m *memberGrpcAdapter) JoinMember(ctx context.Context, req *pb.JoinMember_Request) (*pb.JoinMember_Response, error) {
	log.Println("request: ", req)
	m.memberService.JoinMember()
	return &pb.JoinMember_Response{Code: "200"}, nil
}
