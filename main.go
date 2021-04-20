package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/masterraf21/ecommerce-backend/utils/mysql"
	"github.com/rs/cors"
)

// Server represents server
type Server struct {
	Reader      *sql.DB
	Writer      *sql.DB
	Port        string
	ServerReady chan bool
}

func configureMySQL() (*sql.DB, *sql.DB) {
	readerConfig := mysql.Option{
		Host:     configs.MySQL.ReaderHost,
		Port:     configs.MySQL.ReaderPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.ReaderUser,
		Password: configs.MySQL.ReaderPassword,
	}

	writerConfig := mysql.Option{
		Host:     configs.MySQL.WriterHost,
		Port:     configs.MySQL.WriterPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.WriterUser,
		Password: configs.MySQL.WriterPassword,
	}

	reader, writer, err := mysql.SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}

	log.Println("MySQL connection is successfully established!")

	return reader, writer
}

// Start will start server
func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

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
