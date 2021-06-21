package model

type Product struct {
	Id         int64   `json:"id"`
	CategoryId int64   `json:"categoryId"`
	Images     []*Image   `json:"images"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Rating     float32 `json:"rating"`
	IsSale     bool    `json:"isSale"`
	CreatedAt  int64   `json:"createdAt"`
	ModifiedAt int64   `json:"modifiedAt"`
}
