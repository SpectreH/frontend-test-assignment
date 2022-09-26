package main

import (
	"fmt"
	"log"
	"net/http"
	"test-assgnment/internal/routes"
	"test-assgnment/internal/routes/config"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    config.Port,
		Handler: routes.SetRoutes(),
	}

	fmt.Println("Starting application on port " + config.Port + ". http://localhost" + config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
