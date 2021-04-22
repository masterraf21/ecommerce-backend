package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/apis"
	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/usecases"

	repoMongo "github.com/masterraf21/ecommerce-backend/repositories/mongodb"

	"github.com/masterraf21/ecommerce-backend/utils/mongodb"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rs/cors"
)

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func main() {
	instance := mongodb.ConfigureMongo()
	serverReady := make(chan bool)
	server := Server{
		Instance:    instance,
		Port:        configs.Server.Port,
		ServerReady: serverReady,
	}
	server.Start()
}

// Start will start server
func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

	counterRepo := repoMongo.NewCounterRepo(s.Instance)
	buyerRepo := repoMongo.NewBuyerRepo(s.Instance, counterRepo)
	sellerRepo := repoMongo.NewSellerRepo(s.Instance, counterRepo)
	productRepo := repoMongo.NewProductRepo(s.Instance, counterRepo)
	orderRepo := repoMongo.NewOrderRepo(s.Instance, counterRepo)

	buyerUsecase := usecases.NewBuyerUsecase(buyerRepo)
	sellerUsecase := usecases.NewSellerUsecase(sellerRepo)
	productUsecase := usecases.NewProductUsecase(productRepo, sellerRepo)
	orderUsecase := usecases.NewOrderUsecase(orderRepo, buyerRepo, sellerRepo, productRepo)

	apis.NewBuyerAPI(r, buyerUsecase)
	apis.NewSellerAPI(r, sellerUsecase)
	apis.NewProductAPI(r, productUsecase)
	apis.NewOrderAPI(r, orderUsecase)

	handler := cors.Default().Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port %s!", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Shutting Down Server...")
			log.Fatal(err.Error())
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
