package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	beatspb "github.com/Bamboocho007/go-grps/gen/go/beats/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	beatspb.RegisterBeatsStreamingServiceServer(server, &beatsStreamingServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type beatsStreamingServiceServer struct {
	beatspb.UnimplementedBeatsStreamingServiceServer
}

func (b *beatsStreamingServiceServer) BeatsStream(req *beatspb.BeatsStreamRequest, stream beatspb.BeatsStreamingService_BeatsStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return status.Errorf(codes.Canceled, "stream was canceled")
		default:
			time.Sleep(2 * time.Second)
			err := stream.Send(&beatspb.BeatsStreamResponse{Id: uint32(rand.Intn(100))})

			if err != nil {
				return status.Errorf(codes.Canceled, "stream was canceled")
			}
		}
	}
}
