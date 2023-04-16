package main

import (
	"awesomeProject/api"
	db "awesomeProject/db/sqlc"
	"awesomeProject/util"
	"database/sql"
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
