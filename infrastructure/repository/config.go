package repository

import "go-graphql-boilderplate/infrastructure/db/rdb"

type RepoConfig struct {
	SQLHandler   rdb.SQLHandler
}

