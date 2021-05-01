package usecase

import (
	"go-graphql-boilderplate/entity"
	repository "go-graphql-boilderplate/infrastructure/repository/user"
)

type User struct {
	repo repository.IUser
}

func NewUserService(repo repository.IUser) IUser {
	return &User{repo}
}

func (u *User) GetAll() ([]*entity.User, error) {
	return u.repo.GetAll()
}
