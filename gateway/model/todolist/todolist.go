package todolist

import "time"

type Todolist struct {

	// ID - 事項ID
	ID string

	// Name - 事項名稱
	Name string

	// Status - 事項狀態
	Status bool

	// MemberId - 會員ID
	MemberId string

	// CreatedAt - 建立時間
	CreatedAt time.Time
}

// CreateRequest -
type CreateRequest struct {

	// Name - 待辦事項名稱
	Name string `bson:"name" json:"name"`

	// MemberId - 會員ID
	MemberId string `bson:"memberId" json:"memberId"`
}

// DeleteRequest -
type DeleteRequest struct {

	// ID - 待辦事項ID
	ID string `bson:"id" json:"id"`

	// MemberId - 會員ID
	MemberId string `bson:"memberId" json:"memberId"`
}

// GetRequest -
type GetRequest struct {

	// MemberId - 會員ID
	MemberId string `bson:"memberId" json:"memberId"`
}

// GetResponse -
type GetResponse struct {

	// TodoList - 待辦事項清單
	TodoList []*Todolist
}

// UpdateRequest -
type UpdateRequest struct {

	// ID - 待辦事項ID
	ID string `bson:"id" json:"id"`

	// MemberId - 會員ID
	MemberId string `bson:"memberId" json:"memberId"`
}
