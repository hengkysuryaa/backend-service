package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/repository/web_repo"
	currencyConverterUsecase "github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase/currency_converter"
	orderUsecase "github.com/hengkysuryaa/backend-service/fetch-app/internal/domain/usecase/order"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest/handlers"
	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/cache"
)

func RunRest() {
	restPort := os.Getenv("REST_PORT")
	resourceURL := os.Getenv("RESOURCE_URL")
	currencyConverterURL := os.Getenv("CURRENCY_EXCHANGE_URL")
	currencyConverterAPIKey := os.Getenv("CURRENCY_EXCHANGE_API_KEY")

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// repository
	webRepo := web_repo.New(httpClient, resourceURL, currencyConverterURL, currencyConverterAPIKey)
	mapCache := cache.NewMapCache()

	// order domain
	orderUsecase := orderUsecase.New(webRepo)
	orderHandlers := handlers.NewOrderHandler(orderUsecase)

	// currency converter domain
	currencyConverterUsecase := currencyConverterUsecase.New(webRepo, mapCache)
	currencyConverterHandlers := handlers.NewCurrencyConverterHandler(currencyConverterUsecase)

	r := rest.NewRouter(orderHandlers, currencyConverterHandlers)
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
