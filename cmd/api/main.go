package main

import (
	"log"

	"github.com/hyprhex/blogify/internal/db"
	"github.com/hyprhex/blogify/internal/env"
	"github.com/hyprhex/blogify/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetStr("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetStr("DB_ADDR", "postgres://swap:root@localhost/blogify?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetStr("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	r := app.mount()

	log.Fatal(app.run(r))
}
