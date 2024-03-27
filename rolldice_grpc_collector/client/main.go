package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
    "dice/pb"
)

func main() {
    conn, err := grpc.Dial(
		"server:8080", 
		grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewRollDicerClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req := &pb.RollDiceRequest{}
    resp, err := client.RollDice(ctx, req)
    if err != nil {
        log.Fatalf("could not roll dice: %v", err)
    }

    fmt.Println(resp.DiceNumber)
}
