package srv

import (
	"context"
	"errors"
	"main/member/infra/mongo"
	"main/member/infra/redis"
	"main/member/jwt"

	"golang.org/x/crypto/bcrypt"
)

// SignIn - 登入會員
func SignIn(ctx context.Context, name, password string) (string, int64, error) {

	// 驗證會員
	result, err := mongo.GetByName(ctx, name)
	if err != nil {
		return "", 0, errors.New("註冊失敗，名稱重複")
	}

	// 加密密碼比對
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return "", 0, errors.New("註冊失敗，密碼錯誤")
	}

	// 建立jwt
	jwtToken, err := jwt.GenerateToken(result.ID, result.Name)
	if err != nil {
		return "", 0, errors.New("登入失敗，jwt建立失敗")
	}

	// 寫入redis
	if err := redis.Set(ctx, name, jwtToken); err != nil {
		return "", 0, errors.New("登入失敗，redis寫入失敗")
	}

	return jwtToken, jwt.JWT_TOKEN_LIFE, nil
}
