package model

type HistoryPrice struct {
	Id         int64   `json:"id"`
	ProductId  int64   `json:"productid"`
	OldPrice   float64 `json:"old"`
	ModifiedAt int64   `json:"modifiedAt"`
}
