package srv

import (
	"context"
	"errors"
	member "main/member/infra/mongo"
	"main/member/model"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register - 會員註冊
func Register(ctx context.Context, name, password string) error {

	// 確認名稱是否重複
	_, err := member.GetByName(ctx, name)
	if err == nil {
		return errors.New("註冊失敗，名稱重複")
	}

	// 加密密碼
	hashPassword, err := hash(password)
	if err != nil {
		return errors.New("註冊失敗，密碼加密錯誤")
	}

	// 寫入model
	data := model.Member{
		ID:        bson.NewObjectId().Hex(),
		Name:      name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	// 寫入資料庫
	if err := member.Create(ctx, data); err != nil {
		return errors.New("註冊失敗，寫入資料庫錯誤")
	}

	return nil
}

// hash -  加密
func hash(password string) (string, error) {

	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
