package mongo

import (
	"context"
	"main/member/model"

	"go.mongodb.org/mongo-driver/bson"
)

// GetByName - 透過名稱查詢
func GetByName(ctx context.Context, name string) (model.Member, error) {

	// 連線至collection
	collection := Client.Database("todoList").Collection("member")

	// 查詢資料庫
	var result model.Member
	filter := bson.D{{Key: "name", Value: name}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
