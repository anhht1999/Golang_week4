package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type ProductRepo struct {
	Products map[int64]*model.Product
	autoID   int64
}

var Products ProductRepo 
func init() { 
	Products = ProductRepo{autoID: 0}
	Products.Products = make(map[int64]*model.Product)
	Products.InitData("sql:45312")
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ProductRepo) CreateNewProduct(Product *model.Product) int64 {
	nextID := r.getAutoID() 
	Product.Id = nextID
	r.Products[nextID] = Product 
	return nextID
}

func (r *ProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	// r.CreateNewProduct(&model.Product{
	// 	Title: "Dế Mèn Phiêu Lưu Ký",
	// 	Authors: []model.Author{
	// 		{FullName: "Tô Hoài", Country: "Vietnam"},
	// 		{FullName: "Hames", Country: "Turkey"},
	// 	},
	// 	Rating: 0,
	// })
}

func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {
	return r.Products
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if Product, ok := r.Products[Id]; ok {
		return Product, nil //tìm được
	} else {
		return nil, errors.New("Product not found")
	}
}

func (r *ProductRepo) GetPriceProductById(Id int64) float32 {
	if Product, ok := r.Products[Id]; ok {
		return Product.Price
	}
	return 0
}

func (r *ProductRepo) DeleteProductById(Id int64) error {
	if _, ok := r.Products[Id]; ok {
		delete(r.Products, Id)
		return nil
	} else {
		return errors.New("Product not found")
	}
}

func (r *ProductRepo) UpdateProduct(Product *model.Product) error {
	if _, ok := r.Products[Product.Id]; ok {
		r.Products[Product.Id] = Product
		return nil //tìm được
	} else {
		return errors.New("Product not found")
	}
}

func (r *ProductRepo) Upsert(Product *model.Product) int64 {
	if _, ok := r.Products[Product.Id]; ok {
		r.Products[Product.Id] = Product
		return Product.Id
	} else {
		return r.CreateNewProduct(Product)
	}
}
