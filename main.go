package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/toannguyen3105/techschools_simple_bank/api"
	db "github.com/toannguyen3105/techschools_simple_bank/db/sqlc"
	"github.com/toannguyen3105/techschools_simple_bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connect connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
