package mongo

import (
	"context"
	"log"
	"main/todolist/model"

	"go.mongodb.org/mongo-driver/bson"
)

// GetByMemberId - 透過會員ID查詢
func GetByMemberId(ctx context.Context, id string) ([]*model.Todolist, error) {

	// 連線至collection
	collection := Client.Database("todoList").Collection("todoList")

	// 查詢資料庫
	var results []*model.Todolist
	filter := bson.D{{Key: "memberId", Value: id}}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var result model.Todolist
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)

	return results, nil
}

// GetById - 透過ID查詢
func GetById(ctx context.Context, id string) (*model.Todolist, error) {

	// 連線至collection
	collection := Client.Database("todoList").Collection("todoList")

	// 查詢資料庫
	var result *model.Todolist

	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
