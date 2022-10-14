package main

import (
	"context"
	"fmt"
	"log"

	weatherpb "github.com/Bamboocho007/go-grps/gen/go/weather/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to WeatherStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	weatherStore := weatherpb.NewWeatherStoreServiceClient(conn)
	if _, err := weatherStore.PutWeather(context.Background(), &weatherpb.PutWeatherRequest{
		City:    "New Yourk",
		Country: "USA",
		Degrees: 10,
	}); err != nil {
		return fmt.Errorf("failed to PutWeather: %w", err)
	}

	log.Println("Successfully PutWeather")
	return nil
}
