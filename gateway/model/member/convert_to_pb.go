package member

import pb "main/gateway/protobuf/member"

// RegisterMemberRequestToPb -
func RegisterMemberRequestToPb(in *RegisterMemberRequest) *pb.RegisterRequest {
	if in == nil {
		return new(pb.RegisterRequest)
	}
	return &pb.RegisterRequest{
		Name:     in.Name,
		Password: in.Password,
	}
}

// SignInMemberRequestToPb -
func SignInMemberRequestToPb(in *SignInMemberRequest) *pb.SignInRequest {
	if in == nil {
		return new(pb.SignInRequest)
	}
	return &pb.SignInRequest{
		Name:     in.Name,
		Password: in.Password,
	}
}
