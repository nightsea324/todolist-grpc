package srv

import (
	"context"
	"errors"
	"main/todolist/infra/mongo"
)

// Delete - 刪除事項
func Delete(ctx context.Context, id, memberId string) error {

	// 確認資料是否存在
	result, err := mongo.GetById(ctx, id)
	if err != nil {
		return errors.New("刪除失敗，事項不存在")
	}

	// 確認使用者
	if memberId != result.MemberId {
		return errors.New("刪除失敗，使用者不正確")
	}

	// 刪除資料庫
	if err := mongo.Delete(ctx, id); err != nil {
		return errors.New("刪除失敗，資料庫錯誤")
	}
	return nil
}
