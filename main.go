package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pbatista27/portafolio-serve/db"
	"github.com/pbatista27/portafolio-serve/routers"
)

func main() {

	db.StartUp()

	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	log.Println("servidor corriendo")
	server.ListenAndServe()

}
