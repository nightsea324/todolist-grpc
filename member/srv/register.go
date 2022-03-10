package srv

import (
	"errors"
	"main/model"
	"main/mongo/member"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register - 會員註冊
func Register(name, password string) (string, error) {

	// 接收資料

	// 確認名稱是否重複
	_, err := member.GetByName(name)
	if err == nil {
		return "註冊失敗，名稱重複", errors.New("錯誤")
	}

	// 加密密碼
	hashPassword, err := hash(password)
	if err != nil {
		return "註冊失敗，密碼加密錯誤", err
	}

	// 寫入model
	data := model.Member{
		ID:        bson.NewObjectId().Hex(),
		Name:      name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	// 寫入資料庫
	if err := member.Create(data); err != nil {
		return "註冊失敗，寫入資料庫錯誤", err
	}

	return "註冊成功", nil
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
