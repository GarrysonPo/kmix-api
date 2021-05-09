package main

import (
	"fmt"
	"kmix-api/models"
	"kmix-api/routers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"),
	)
	models.InitDB(dataSourceName)
	defer models.CloseDB()

	router := routers.SetupRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Name:\t\tkmix-api")
	fmt.Println("Version:\tv1")
	fmt.Println("Addres:\t\thttp://localhost:8000/api/v1")
	fmt.Println("Status:\t\tstarted")
	log.Fatal(srv.ListenAndServe())
}
