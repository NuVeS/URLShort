package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/NuVeS/URLShort/cmd/apis/auth"
	"github.com/NuVeS/URLShort/cmd/apis/urls"
	"github.com/NuVeS/URLShort/cmd/router"
	"github.com/NuVeS/URLShort/cmd/storage"
)

func main() {
	storage := &storage.StorageDB{}

	router := router.NewRouter()

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	auth.DB = storage
	urls.DB = storage

	go srv.ListenAndServe()

	fmt.Println()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
