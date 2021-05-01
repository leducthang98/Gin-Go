package repository

import (
	"go-graphql-boilderplate/entity"
	"go-graphql-boilderplate/infrastructure/db/rdb"
	"go-graphql-boilderplate/infrastructure/repository"
)

type User struct {
	sqlHandler rdb.SQLHandler
}

func NewUserRepository(config repository.RepoConfig) IUser {
	userRepository := new(User)
	userRepository.sqlHandler = config.SQLHandler
	return userRepository
}

func (u *User) GetAll() (users []*entity.User, err error) {
	sql := "select * from user"
	err = u.sqlHandler.Query(&users, sql)
	return
}
