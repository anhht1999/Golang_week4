package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type CategoryRepo struct {
	Categorys map[int64]*model.Category
	autoID    int64 //đây là biến đếm tự tăng gán giá trị cho id của Category
}

var Categorys CategoryRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Categorys = CategoryRepo{autoID: 0}
	Categorys.Categorys = make(map[int64]*model.Category)
	Categorys.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *CategoryRepo
func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CategoryRepo) CreateNewCategory(Category *model.Category) int64 {
	nextID := r.getAutoID() 
	Category.Id = nextID
	r.Categorys[nextID] = Category 
	return nextID
}

func (r *CategoryRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)
}

func (r *CategoryRepo) GetAllCategorys() map[int64]*model.Category {
	return r.Categorys
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*model.Category, error) {
	if Category, ok := r.Categorys[Id]; ok {
		return Category, nil //tìm được
	} else {
		return nil, errors.New("Category not found")
	}
}

func (r *CategoryRepo) DeleteCategoryById(Id int64) error {
	if _, ok := r.Categorys[Id]; ok {
		delete(r.Categorys, Id)
		return nil
	} else {
		return errors.New("Category not found")
	}
}

func (r *CategoryRepo) UpdateCategory(Category *model.Category) error {
	if _, ok := r.Categorys[Category.Id]; ok {
		r.Categorys[Category.Id] = Category
		return nil //tìm được
	} else {
		return errors.New("Category not found")
	}
}

func (r *CategoryRepo) Upsert(Category *model.Category) int64 {
	if _, ok := r.Categorys[Category.Id]; ok {
		r.Categorys[Category.Id] = Category
		return Category.Id
	} else {
		return r.CreateNewCategory(Category)
	}
}
