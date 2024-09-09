package main

import (
	"log"

	"github.com/hyprhex/blogify/internal/db"
	"github.com/hyprhex/blogify/internal/env"
	"github.com/hyprhex/blogify/internal/store"
)

func main() {
	addr := env.GetStr("DB_ADDR", "postgres://swap:root@localhost/blogify?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
