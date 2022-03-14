package member

// RegisterMemberRequest -
type RegisterMemberRequest struct {

	// Name - 會員名稱
	Name string `bson:"name" json:"name"`

	// Password - 會員密碼
	Password string `bson:"password" json:"password"`
}

// SignInMemberRequest -
type SignInMemberRequest struct {

	// Name - 會員名稱
	Name string `bson:"name" json:"name"`

	// Password - 會員密碼
	Password string `bson:"password" json:"password"`
}
