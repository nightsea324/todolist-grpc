package mongo

import (
	"context"
	"main/todolist/model"
)

// Create - 新增待辦事項
func Create(ctx context.Context, todoList model.Todolist) error {

	// 連線至collection
	collection := Client.Database("todoList").Collection("todoList")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, todoList)
	if err != nil {
		return err
	}

	return nil
}
