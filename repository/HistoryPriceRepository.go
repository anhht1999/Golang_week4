package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type HistoryPriceRepo struct {
	HistoryPrices map[int64]*model.HistoryPrice
	autoID        int64
}

var HistoryPrice HistoryPriceRepo

func init() {
	HistoryPrice = HistoryPriceRepo{autoID: 0}
	HistoryPrice.HistoryPrices = make(map[int64]*model.HistoryPrice)
	HistoryPrice.InitData("sql:45312")
}

func (r *HistoryPriceRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *HistoryPriceRepo) CreateNewHistoryPrice(HistoryPrice *model.HistoryPrice) error {
	nextID := r.getAutoID()
	HistoryPrice.Id = nextID
	r.HistoryPrices[nextID] = HistoryPrice
	return nil
}
func (r *HistoryPriceRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewHistoryPrice(&model.HistoryPrice{
		Id:        1,
		ProductId: 1,
		OldPrice:  10,
		ModifiedAt:1615976361000,
	})

}

func (r *HistoryPriceRepo) GetAllHistoryPrices() map[int64]*model.HistoryPrice {
	return r.HistoryPrices
}

//del
func (r *HistoryPriceRepo) DeleteHistoryPriceById(id int64) error {
	if _, ok := r.HistoryPrices[id]; ok {
		delete(r.HistoryPrices, id)
		return nil
	} else {
		return errors.New("HistoryPrice not found")
	}
}
