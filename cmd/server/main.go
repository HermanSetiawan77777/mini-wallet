package main

import (
	"context"
	"herman-technical-julo/config"
	"herman-technical-julo/internal/httpserver"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.LoadEnv()

	dbs, err := buildDatabases()
	if err != nil {
		panic("Failed to connect JULO DB")
	}
	appContainer := setupAppContainer(dbs)
	appServer := httpserver.NewServer(config.Port(), appContainer)

	go func() {
		err := appServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Server stopped: %s\n", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := appServer.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server: %s", err.Error())
	}

	log.Println("Server gracefully shutdown")

}
