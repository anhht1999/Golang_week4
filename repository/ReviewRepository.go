package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type ReviewRepo struct {
	reviews  map[int64]*model.Review
	autoID int64 
}

var Reviews ReviewRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45314")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *ProductRepo
func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	review.Id = nextID
	r.reviews[nextID] = review //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}


func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	// r.CreateNewReview(&model.Review{
	// 	ProductId: 1,
	// 	Comment: "Hay qua",
	// 	Rating: 4,
	// })

	// r.CreateNewReview(&model.Review{
	// 	ProductId: 1,
	// 	Comment: "Hay qua 1",
	// 	Rating: 5,
	// })

	// r.CreateNewReview(&model.Review{
	// 	ProductId: 2,
	// 	Comment: "Hay qua 3",
	// 	Rating: 4,
	// })

	// r.CreateNewReview(&model.Review{
	// 	ProductId: 2,
	// 	Comment: "Hay qua",
	// 	Rating: 5,
	// })
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) FindReviewById(Id int64) (*model.Review, error) {

	
	if review, ok := r.reviews[Id]; ok {
		return review, nil //tìm được
	} else {
		return nil, errors.New("review not found")
	}
}

func (r *ReviewRepo) AverageRating(ProductId int64) (result map[int64]float32) {

	sumRating := make(map[int64]int)
	numberRating := make(map[int64]int)
	result = make(map[int64]float32)

	for _, value := range r.reviews {
		numberRating[value.ProductId]++
		sumRating[value.ProductId] += value.Rating
	}
	for key := range numberRating {
		result[key] = float32(sumRating[key]) / float32(numberRating[key])
	}
	return result
}

func (r *ReviewRepo) DeleteReview(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpdateReview(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil //tìm được
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpsertReview(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}


