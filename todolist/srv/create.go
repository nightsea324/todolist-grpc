package srv

import (
	"context"
	"errors"
	"main/todolist/infra/mongo"
	"main/todolist/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Create - 新增事項
func Create(ctx context.Context, name, memberId string) error {

	// 寫入model
	todoList := model.Todolist{
		ID:        bson.NewObjectId().Hex(),
		Name:      name,
		Status:    false,
		MemberId:  memberId,
		CreatedAt: time.Now(),
	}

	// 寫入資料庫
	if err := mongo.Create(ctx, todoList); err != nil {
		return errors.New("新增失敗，資料庫錯誤")
	}
	return nil
}
