package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/menemsalama/simplenote-go/api"
	"github.com/menemsalama/simplenote-go/internal/database"
	"github.com/menemsalama/simplenote-go/migrations"
)

func main() {
	// dbURL := fmt.Sprintf(
	// 	"host=%s port=%s user=%s dbname=%s sslmode=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_SSL"),
	// )
	dbURL := os.Getenv("DATABASE_URL")

	if err := database.ConnectToPG(dbURL); err != nil {
		log.Fatal(err)
	}
	defer database.PG.Close()

	database.PG.LogMode(true)

	migrations.Run()

	api := api.NewAPI()

	log.Println("Starting.. port :8080")

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Fatal(err)
	}
}
