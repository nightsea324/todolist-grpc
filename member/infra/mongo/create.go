package mongo

import (
	"context"
	"main/member/model"
)

// Create - 新增會員
func Create(ctx context.Context, member model.Member) error {

	// 連線至collection
	collection := Client.Database("todoList").Collection("member")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, member)
	if err != nil {
		return err
	}

	return nil
}
