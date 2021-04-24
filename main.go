package main

import (
	"fmt"
	"kmix-api/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	router := routers.SetupRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Project:\tkmix-api")
	fmt.Println("Version:\tv1")
	fmt.Println("Addres:\t\thttp://localhost:8000/api/v1")
	fmt.Println("Status:\t\tstarted")
	log.Fatal(srv.ListenAndServe())
}
