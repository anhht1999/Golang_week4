package model

type Image struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"ProductId"`
	Url       string `json:"url"`
}
