package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type ImageRepo struct {
	Images  map[int64]*model.Image
	autoID int64 
}

var Images ImageRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Images = ImageRepo{autoID: 0}
	Images.Images = make(map[int64]*model.Image)
	Images.InitData("sql:45314")
}

func (r *ImageRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ImageRepo) CreateNewImage(Image *model.Image) map[int64]*model.Image {
	nextID := r.getAutoID() 
	Image.Id = nextID
	r.Images[nextID] = Image
	return r.Images
}


func (r *ImageRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)
}

func (r *ImageRepo) GetAllImages() map[int64]*model.Image {
	return r.Images
}

func (r *ImageRepo) FindImageById(Id int64) (*model.Image, error) {

	
	if Image, ok := r.Images[Id]; ok {
		return Image, nil //tìm được
	} else {
		return nil, errors.New("Image not found")
	}
}

func (r *ImageRepo) DeleteImage(Id int64) error {
	if _, ok := r.Images[Id]; ok {
		delete(r.Images, Id)
		return nil
	} else {
		return errors.New("Image not found")
	}
}

func (r *ImageRepo) UpdateImage(Image *model.Image) error {
	if _, ok := r.Images[Image.Id]; ok {
		r.Images[Image.Id] = Image
		return nil //tìm được
	} else {
		return errors.New("Image not found")
	}
}

// func (r *ImageRepo) UpsertImage(Image *model.Image) int64 {
// 	if _, ok := r.Images[Image.Id]; ok {
// 		r.Images[Image.Id] = Image
// 		return Image.Id
// 	} else {
// 		return r.CreateNewImage(Image)
// 	}
// }


