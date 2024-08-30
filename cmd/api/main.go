package main

import (
	"log"

	"github.com/hyprhex/blogify/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetStr("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	r := app.mount()

	log.Fatal(app.run(r))
}
