package main

import (
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/seba5dev/hormigasto-backend/internal/database"
	"github.com/seba5dev/hormigasto-backend/models"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	database.Connect()

	// Migrate the database schema
	// This will create the tables if they do not exist
	// and update the schema if necessary
	err = database.DB.AutoMigrate(
		&models.User{},
		&models.AccountType{},
		&models.Account{},
		&models.TransactionType{},
		&models.TransactionCategory{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", r)
}
