package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

var connectionError error

func timer(name string) error {
	var wg sync.WaitGroup
	for j := 0; j < 5; j++ {
		start := time.Now()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				var titleInput = fmt.Sprintf("GoLang %v", i)
				if err := createDispositivo(titleInput); err != nil {
					log.Println(err)
					return
				}
			}(i)
		}
		wg.Wait()
		fmt.Printf("iteration %v on %s took %v\n", j, name, time.Since(start))
	}
	return nil

}

func main() {
	// Open database connection:
	database, connectionError = sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}
	defer database.Close()

	// Connection configs:
	database.SetMaxIdleConns(15)
	database.SetMaxOpenConns(15)

	// Execute timer
	if err := timer(fmt.Sprintf("main")); err != nil {
		panic(err)
	}
}
