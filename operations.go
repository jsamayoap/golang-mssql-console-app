package main

import (
	"context"
	"fmt"
	"time"
)

func createDispositivo(titleInput string) error {
	var err error
	ctx, cancelfunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelfunc()

	queryStatement := fmt.Sprintf("INSERT INTO  ( [nombre] ) VALUES ( '%v' );", titleInput)

	query, err := database.Prepare(queryStatement)
	if err != nil {
		fmt.Printf("Query err: %v", err)
	}

	_, queryErr := query.QueryContext(ctx)

	if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}

	return nil
}
