package repository

import (
	"errors"
	"fmt"

	"github.com/anhht1999/Golang_week4/model"
)

type UserRepo struct {
	Users  map[int64]*model.User
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của User
}

var Users UserRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Users = UserRepo{autoID: 0}
	Users.Users = make(map[int64]*model.User)
	Users.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *UserRepo
func (r *UserRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *UserRepo) CreateNewUser(User *model.User) int64 {
	nextID := r.getAutoID() 
	User.Id = nextID
	r.Users[nextID] = User
	return nextID
}

func (r *UserRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

}

func (r *UserRepo) GetAllUsers() map[int64]*model.User {
	return r.Users
}

func (r *UserRepo) FindUserById(Id int64) (*model.User, error) {
	if User, ok := r.Users[Id]; ok {
		return User, nil //tìm được
	} else {
		return nil, errors.New("User not found")
	}
}

func (r *UserRepo) DeleteUserById(Id int64) error {
	if _, ok := r.Users[Id]; ok {
		delete(r.Users, Id)
		return nil
	} else {
		return errors.New("User not found")
	}
}

func (r *UserRepo) UpdateUser(User *model.User) error {
	if _, ok := r.Users[User.Id]; ok {
		r.Users[User.Id] = User
		return nil //tìm được
	} else {
		return errors.New("User not found")
	}
}

func (r *UserRepo) Upsert(User *model.User) int64 {
	if _, ok := r.Users[User.Id]; ok {
		r.Users[User.Id] = User
		return User.Id
	} else {
		return r.CreateNewUser(User)
	}
}
