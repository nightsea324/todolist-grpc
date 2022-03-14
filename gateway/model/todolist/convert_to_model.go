package todolist

import (
	pb "main/gateway/protobuf/todolist"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// ConvertTodoListToModel -
func ConvertTodoListToModel(in []*pb.Todolist) []*Todolist {
	result := make([]*Todolist, len(in))

	if len(in) == 0 {
		return result
	}

	for i, v := range in {
		result[i] = todoToModel(v)
	}
	return result
}

// todoToModel -
func todoToModel(in *pb.Todolist) *Todolist {
	return &Todolist{
		ID:       in.ID,
		Name:     in.Name,
		Status:   in.Status,
		MemberId: in.MemberId,
		CreatedAt: func() time.Time {
			t, _ := ptypes.Timestamp(in.CreatedAt)
			return t
		}(),
	}
}
