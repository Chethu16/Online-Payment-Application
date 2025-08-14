package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func start(conn *connection) {
	router := mux.NewRouter()
	server := &http.Server{
		Addr: os.Getenv("SERVER_PORT"),
		Handler: router,
	}
	// All Routes Here
	//authenticationRoutes(router,conn)


	log.Println("server is running on port ",os.Getenv("SERVER_PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("unable to start the server:",err.Error())
	}
}

//func authenticationRoutes(router *mux.Router , conn *connection) {
	//router.HandleFunc("/login" , )
	//router.HandleFunc("/register" ,)
//}