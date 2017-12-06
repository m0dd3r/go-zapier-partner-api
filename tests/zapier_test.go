package tests

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/m0dd3r/go-zapier-partner-api/zapier"
)

var (
	client *zapier.Client
	ctx    context.Context

	auth bool
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	version := os.Getenv("ZAPIER_PARTNER_API_VERSION")
	if version == "" {
		panic("CAN'T RUN TESTS WITHOUT VERSION")
	}

	client_id := os.Getenv("ZAPIER_PARTNER_API_CLIENT_ID")
	if client_id == "" {
		fmt.Printf("NO CLIENT_ID, SOME TESTS MAY NOT RUN")
	} else {
		auth = true
	}
	ctx = context.Background()
	client = zapier.NewClient(version, client_id, nil)
	fmt.Println("TEST CLIENT INITIALIZED")
}
