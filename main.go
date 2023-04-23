package main

import (
	"database/sql"
	"github.com/Reoptima/GoLearnPrject/api"
	db "github.com/Reoptima/GoLearnPrject/db/sqlc"
	_ "github.com/Reoptima/GoLearnPrject/docs"
	"github.com/Reoptima/GoLearnPrject/util"
	_ "github.com/lib/pq"
	"log"
)

//	@title			CRM API Server
//	@version		1.3
//	@description	API Сервер для совершения транзакций внутри CRM

//	@host	localhost:8083
//	@Base	path /

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
