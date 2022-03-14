package todolist

import pb "main/gateway/protobuf/todolist"

// CreateRequestToPb -
func CreateRequestToPb(in *CreateRequest) *pb.CreateRequest {
	if in == nil {
		return new(pb.CreateRequest)
	}
	return &pb.CreateRequest{
		Name:     in.Name,
		MemberId: in.MemberId,
	}
}

// DeleteRequestToPb -
func DeleteRequestToPb(in *DeleteRequest) *pb.DeleteRequest {
	if in == nil {
		return new(pb.DeleteRequest)
	}
	return &pb.DeleteRequest{
		ID:       in.ID,
		MemberId: in.MemberId,
	}
}

// GetRequestToPb -
func GetRequestToPb(in *GetRequest) *pb.GetRequest {
	if in == nil {
		return new(pb.GetRequest)
	}
	return &pb.GetRequest{
		MemberId: in.MemberId,
	}
}

// UpdateRequestToPb -
func UpdateRequestToPb(in *UpdateRequest) *pb.UpdateRequest {
	if in == nil {
		return new(pb.UpdateRequest)
	}
	return &pb.UpdateRequest{
		ID:       in.ID,
		MemberId: in.MemberId,
	}
}
