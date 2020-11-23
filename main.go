package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/seijihg/api_template_mongodb/controllers"
	"github.com/seijihg/api_template_mongodb/database"
)

func main() {

	client := database.ConnectDB()
	golangDB := client.Database("golang")

	// Prevent leaking.
	defer client.Disconnect(context.Background())

	// Routing
	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	apiV1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API v1")
	})
	apiV1.HandleFunc("/user", controllers.CreateUser(golangDB)).Methods("POST")

	// Start server
	srv := &http.Server{
		Handler: router,
		Addr:    ":3000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting up...")
	log.Fatal(srv.ListenAndServe())
}
