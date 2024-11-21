package main

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	_, err := bigquery.NewClient(
		ctx,
		"foo",
		option.WithEndpoint("http://localhost:8888"),
		option.WithoutAuthentication(),
	)
	if err != nil {
		log.Fatalf("failed to bigquery.NewClient: %v", err)
	}
}
