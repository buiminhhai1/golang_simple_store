package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/thuhangpham/simplestores/api"
	db "github.com/thuhangpham/simplestores/db/sqlc"
	"github.com/thuhangpham/simplestores/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connct to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(*store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
