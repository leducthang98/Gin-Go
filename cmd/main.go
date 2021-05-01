package main

import (
	"context"
	server "go-graphql-boilderplate/api"
	"go-graphql-boilderplate/config"
)

type User struct {
	Id                 int
	Username, Password string
}

func main() {
	srvCfg := new(config.Server)

	// Fix tạm thôi đừng chửi tao, địt mẹ mày Thiên ạ
	srvCfg.Port = 8080
	srvCfg.Mysql.DBName = "default"
	srvCfg.Mysql.Port = 3306
	srvCfg.Mysql.Host = "localhost"
	srvCfg.Mysql.User = "root"
	srvCfg.Mysql.Password = "codedidungso.me"

	srv := server.New(srvCfg)

	if err := srv.Start(context.Background()); err != nil {
		panic(err)
	}

	defer srv.Close()
}
