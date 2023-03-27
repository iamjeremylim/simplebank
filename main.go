package main

import (
	"database/sql"
	"log"

	"github.com/iamjeremylim/simplebank/util"

	"github.com/iamjeremylim/simplebank/api"
	db "github.com/iamjeremylim/simplebank/db/sqlc"
	_ "github.com/lib/pq" // for server to connect with db
)

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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
