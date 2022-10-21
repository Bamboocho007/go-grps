package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	beatspb "github.com/Bamboocho007/go-grps/gen/go/beats/v1"
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
		return fmt.Errorf("failed to connect to StreamingStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)
	beatsService := beatspb.NewBeatsStreamingServiceClient(conn)
	stream, err := beatsService.BeatsStream(context.Background(), &beatspb.BeatsStreamRequest{Id: uint32(1)})
	if err != nil {
		fmt.Errorf("failed to stream: %w", err)
	}

	counter := 0

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}
		counter++
		log.Println("Successfully stream", res.Id)

		if counter > 5 {
			stream.SendMsg(errors.New("close!!!"))
		}
	}
}
