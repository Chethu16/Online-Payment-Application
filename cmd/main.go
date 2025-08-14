package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("unable to load env file:", err.Error())
	}

	if os.Getenv("MODE") == "dev" {
		log.Println("server is running in development mode.")
	}

	if os.Getenv("MODE") == "prod" {
		log.Println("server is running in production mode.")
	}
}

func main() {
	conn := connectToDatabase()
	conn.checkDatabaseStatus()
	defer conn.closeDatabaseConnection()
	start(conn)
}
