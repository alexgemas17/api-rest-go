package main

import (
	"log"

	"github.com/alexgemas17/api-rest-go/api"
	"github.com/alexgemas17/api-rest-go/api/envs"
	"github.com/alexgemas17/api-rest-go/api/store"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 envs.Envs.DBUser,
		Passwd:               envs.Envs.DBPassword,
		Addr:                 envs.Envs.DBAddress,
		DBName:               envs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlRepository := store.NewSQLRepository(cfg)

	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal((err))
	}

	repository := store.NewStore(db)

	api := api.NewApiServer(":3000", repository)
	api.Serve()
}
