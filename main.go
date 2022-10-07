package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("starting program")
	databaseUrl := "postgres://postgres:postgres@localhost:55003"
	dbPool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	ExecuteSelectQuery(dbPool)
	ExecuteFunction(dbPool, 7)
	log.Println("stopping program")
}
func ExecuteSelectQuery(dbPool *pgxpool.Pool) {
	log.Println("starting execution of select query")

	// execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from ttt")
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		// convert DB types to Go types
		id := values[0].(int32)
		firstName := values[1].(string)
		lastName := values[2].(string)
		dateOfBirth := values[3].(time.Time)
		log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}

}

func ExecuteFunction(dbPool *pgxpool.Pool, id int) {
	log.Println("starting execution of database function")

	// execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from get_ttt($1)", id)
	log.Println("input id: ", id)
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		//convert DB types to Go types
		firstName := values[0].(string)
		lastName := values[1].(string)
		dateOfBirth := values[2].(time.Time)

		log.Println("[first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}

}
