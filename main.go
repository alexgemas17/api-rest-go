package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlRepository := NewSQLRepository(cfg)

	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal((err))
	}

	repository := NewStore(db)

	api := NewApiServer(":3000", repository)
	api.Serve()
}
