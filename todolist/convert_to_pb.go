package todolist

import (
	"main/todolist/model"
	todolist "main/todolist/protobuf"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ConverTodoListToPb - 轉換model至pb
func ConverTodoListToPb(in []*model.Todolist) []*todolist.Todolist {
	result := make([]*todolist.Todolist, len(in))
	if len(in) == 0 {
		return result
	}
	for i, v := range in {
		result[i] = &todolist.Todolist{
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
