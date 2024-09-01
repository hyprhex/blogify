package main

import (
	"log"

	"github.com/hyprhex/blogify/internal/env"
	"github.com/hyprhex/blogify/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetStr("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	r := app.mount()

	log.Fatal(app.run(r))
}
