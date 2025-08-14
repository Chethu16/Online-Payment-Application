package main

import (
	"database/sql"
	"log"
	"os"

	_"github.com/lib/pq"
)

type connection struct {
	db *sql.DB
}

func connectToDatabase() *connection {
	conn , err := sql.Open(os.Getenv("DATABASE_NAME") , os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("unable to connect to database:",err.Error())
	}
	log.Println("database connection successfull")
	return &connection{
		conn,
	}
}

func (c *connection) checkDatabaseStatus() {
	if err := c.db.Ping(); err != nil {
		log.Fatalln("database not working properly:",err.Error())
	}
	log.Println("database working properly")
}

func (c *connection) closeDatabaseConnection() {
	if err := c.db.Close(); err != nil {
		log.Fatalln("database closed with an error:",err.Error())
	}
	log.Println("database connection closed successfully")
}