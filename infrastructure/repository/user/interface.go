package repository

import (
	"go-graphql-boilderplate/entity"
)

type IUser interface {
	GetAll() ([]*entity.User, error)
}

