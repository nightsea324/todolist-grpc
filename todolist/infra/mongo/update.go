package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update - 完成待辦事項
func Update(ctx context.Context, id string) error {

	// 連線至collection
	collection := Client.Database("todoList").Collection("todoList")

	// 更新資料庫資料
	filter := bson.D{{Key: "id", Value: id}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.D{
			{Key: "status", Value: true}},
		}}
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
