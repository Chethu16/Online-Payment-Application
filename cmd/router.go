package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Chethu16/Online-Payment-Application/internal/handlers"
	"github.com/Chethu16/Online-Payment-Application/internal/repositories"
	"github.com/Chethu16/Online-Payment-Application/pkg/queries"
	"github.com/gorilla/mux"
)

func start(conn *connection) {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: router,
	}

	query := queries.NewQuery(conn.db)
	query.InitDatabase()
	// All Routes Here
	authenticationRoutes(router, conn)

	log.Println("server is running on port ", os.Getenv("SERVER_PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("unable to start the server:", err.Error())
	}
}

func authenticationRoutes(router *mux.Router, conn *connection) {
	repo := repositories.NewAuthenticationRepository(conn.db)
	handler := handlers.NewAuthenticationHandler(repo)
	router.HandleFunc("/login", handler.LoginRequest).Methods("POST")
	router.HandleFunc("/register", handler.RegisterRequest).Methods("POST")
}
