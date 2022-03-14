package todolist

import (
	"main/todolist/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ConvertTodoListToPb - 轉換model至pb
func ConvertTodoListToPb(in []*model.Todolist) []*Todolist {
	result := make([]*Todolist, len(in))
	if len(in) == 0 {
		return result
	}
	for i, v := range in {
		result[i] = &Todolist{
			ID:       v.ID,
			Name:     v.Name,
			Status:   v.Status,
			MemberId: v.MemberId,
			CreatedAt: func() *timestamppb.Timestamp {
				t, _ := ptypes.TimestampProto(v.CreatedAt)
				return t
			}(),
		}
	}
	return result
}
