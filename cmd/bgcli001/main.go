package main

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

func main() {
	ctx := context.Background()
	_, err := bigquery.NewClient(ctx, "foo")
	if err != nil {
		log.Fatalf("failed to bigquery.NewClient: %v", err)
	}
}
