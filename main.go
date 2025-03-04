package main

import (
	"comprezo/config"
	"comprezo/router"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fConfig := flag.String("config", "config.json", "path to config.json")
	flag.Parse()

	cfg, err := config.Load(*fConfig)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Init(cfg),
	}

	log.Println("Server is running on", cfg.Port)
	return srv.ListenAndServe()
}
