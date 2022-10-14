package main

import (
	"context"
	"fmt"
	"log"
	"net"

	weatherpb "github.com/Bamboocho007/go-grps/gen/go/weather/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	weatherpb.RegisterWeatherStoreServiceServer(server, &weatherStoreServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type weatherStoreServiceServer struct {
	weatherpb.UnimplementedWeatherStoreServiceServer
}

func (s *weatherStoreServiceServer) GetWeather(ctx context.Context, req *weatherpb.GetWeatherRequest) (*weatherpb.GetWeatherResponse, error) {
	city := req.GetCity()
	country := req.GetCountry()
	log.Println("Got a request to get a weather for city", city, "for country", country)

	return &weatherpb.GetWeatherResponse{}, nil
}

func (s *weatherStoreServiceServer) PutWeather(ctx context.Context, req *weatherpb.PutWeatherRequest) (*weatherpb.PutWeatherResponse, error) {
	city := req.GetCity()
	country := req.GetCountry()
	log.Println("Got a request to create a weather for city", city, "for country", country)

	return &weatherpb.PutWeatherResponse{}, nil
}

func (s *weatherStoreServiceServer) DeleteWeather(ctx context.Context, req *weatherpb.DeleteWeatherRequest) (*weatherpb.DeleteWeatherResponse, error) {
	id := req.Id
	log.Println("Got a request to delete a weather with id", id)

	return &weatherpb.DeleteWeatherResponse{}, nil
}
