package srv

import (
	"context"
	"errors"
	"main/todolist/infra/mongo"
)

// Update - 完成事項
func Update(ctx context.Context, id, memberId string) error {

	// 確認資料是否存在
	result, err := mongo.GetById(ctx, id)

	if err != nil {
		return errors.New("更新失敗，待辦事項不存在")
	}

	// 確認使用者是否正確
	if memberId != result.MemberId {
		return errors.New("更新失敗，使用者錯誤")
	}

	// 更新資料庫資料
	if err := mongo.Update(ctx, id); err != nil {
		return errors.New("更新失敗，資料庫錯誤")
	}

	return nil
}
