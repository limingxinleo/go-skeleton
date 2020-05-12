package dao

import (
	"app/kernel/provider"
	"app/model"
)

type UserDao struct {
}

func (this UserDao) First(id int) model.User {
	var user model.User
	provider.DB.Where("id = ?", id).First(&user)
	return user
}
