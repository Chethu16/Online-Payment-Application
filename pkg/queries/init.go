package queries

import (
	"database/sql"
	"log"
)

type query struct {
	db *sql.DB
}

func NewQuery(db *sql.DB) *query {
	return &query{
		db,
	}
}

func (q *query) InitDatabase() {
	queries := []string{
		`
			CREATE TABLE IF NOT EXISTS users(
				user_id VARCHAR PRIMARY KEY,
				user_name VARCHAR NOT NULL,
				user_email VARCHAR NOT NULL UNIQUE,
				user_password VARCHAR NOT NULL
			)
		`,
	}
	tx, err := q.db.Begin()

	if err != nil {
		tx.Rollback()
		log.Fatalln("unable to initilize database:", err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, query := range queries {
		_, err = tx.Exec(query)
		if err != nil {
			log.Fatalln("unable to execuite query:", err.Error())
		}
	}
	log.Println("query execuited successfully")
}
