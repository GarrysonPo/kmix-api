package main_test

import (
	"fmt"
	"kmix-api/models"
	"kmix-api/routers"
	"log"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

func TestMain(m *testing.M) {

	err := godotenv.Load("../.env")
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

	router = routers.SetupRouter()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
