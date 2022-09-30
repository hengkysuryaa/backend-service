package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest"
)

func RunRest() {
	restPort := os.Getenv("REST_PORT")

	r := rest.NewRouter()
	s := http.Server{
		Addr:    restPort,
		Handler: r,
	}

	// Graceful shutdown
	signalChan := make(chan os.Signal, 1)
	cleanupChan := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan

		log.Println("receive an interrupt. starting shutdown")

		err := s.Shutdown(context.Background())
		if err != nil {
			defer os.Exit(1)
			log.Println("shutdown error")
			return
		}

		cleanupChan <- true
	}()

	log.Println("Rest server started on", restPort, "port")

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		defer os.Exit(1)
		log.Println("ListenAndServe error")
		return
	}

	<-cleanupChan

	log.Println("shutdown completed")
}
