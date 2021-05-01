package ioc

import (
	"go-graphql-boilderplate/infrastructure/db/rdb"
	"go-graphql-boilderplate/infrastructure/repository"
	repository2 "go-graphql-boilderplate/infrastructure/repository/user"
	usecase "go-graphql-boilderplate/usecase/user"
)

type Resolver struct {
	UserService usecase.IUser
}

type ResolverCfg struct {
	SQLHandler rdb.SQLHandler
}

func NewResolver(resolverCfg ResolverCfg) Resolver {
	repoConfig := repository.RepoConfig{
		SQLHandler: resolverCfg.SQLHandler,
	}
	r := Resolver{
		UserService: usecase.NewUserService(repository2.NewUserRepository(repoConfig)),
	}

	return r
}
