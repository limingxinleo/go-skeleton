package dao

import "app/model"

type UserDao struct {
}

func (this UserDao) First(id int) model.User {
	var user model.User
	DB.Where("id = ?", id).First(user)
	return user
}
