package srv

import (
	"context"
	"errors"
	"main/todolist/infra/mongo"
	"main/todolist/model"
)

// Get - 取得事項
func Get(ctx context.Context, memberId string) ([]*model.Todolist, error) {

	// 透過會員ID獲取資料
	results, err := mongo.GetByMemberId(ctx, memberId)
	if err != nil {
		return results, errors.New("資料庫錯誤")
	}

	return results, nil
}
