package model

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Birthday   string `json:"birthday"`
	Status     bool   `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	ModifiedAt int64  `json:"modifiedAt"`
}
