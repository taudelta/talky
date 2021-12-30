package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	talky "github.com/taudelta/talky/internal"
	"github.com/taudelta/talky/internal/config"
	"github.com/taudelta/talky/internal/storage"
)

func main() {

	envConfig := config.InitConfig()

	pgStorage := storage.NewPostgreSQLStorage(envConfig)

	phraseRepo := storage.NewPhraseRepository()

	router := mux.NewRouter()

	talky.InitApiV1(router, talky.Dependencies{
		Database:   pgStorage,
		PhraseRepo: phraseRepo,
	})

	srv := &http.Server{
		Addr:    envConfig.Addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("server listen error", err)
			return
		}
	}()

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	log.Println("Start server")
	sig := <-sigCh
	log.Println("signal received", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	log.Println("server stopped")
}
