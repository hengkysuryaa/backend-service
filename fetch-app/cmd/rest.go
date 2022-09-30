package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository/web_repo"
	orderUsecase "github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase/order"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest"
	orderHandler "github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest/handlers"
)

func RunRest() {
	restPort := os.Getenv("REST_PORT")
	resourceURL := os.Getenv("RESOURCE_URL")
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	webRepo := web_repo.New(httpClient, resourceURL)
	orderUsecase := orderUsecase.New(webRepo)
	orderHandlers := orderHandler.NewOrderHandler(orderUsecase)

	r := rest.NewRouter(orderHandlers)
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
